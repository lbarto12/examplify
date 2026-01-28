package core

import (
	"context"
	"encoding/json"
	"fmt"
	"server/api/tools/features/imageanalysis"
	"server/sqlc/sqlgen"
	"strings"

	"github.com/google/uuid"
	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/responses"
)

func (core Core) AnalyzeCollection(
	ctx context.Context,
	userID uuid.UUID,
	collectionID uuid.UUID,
	kind sqlgen.AnalysisType,
) (*CollectionAnalysis, error) {

	q := core.Queries

	// Auth check
	if _, err := q.GetCollection(ctx, sqlgen.GetCollectionParams{
		UserID: userID,
		ID:     collectionID,
	}); err != nil {
		return nil, err
	}

	// Ensure text exists
	if err := core.ensureExtractions(ctx, userID, collectionID); err != nil {
		return nil, err
	}

	// Snapshot
	snapshot, err := core.createSnapshot(ctx, collectionID)
	if err != nil {
		return nil, err
	}

	// Run AI
	result, err := core.runAnalysis(ctx, snapshot.CombinedContent, kind)
	if err != nil {
		return nil, err
	}

	row, err := q.CreateCollectionAnalysis(ctx, sqlgen.CreateCollectionAnalysisParams{
		SnapshotID: snapshot.ID,
		Type:       sqlgen.AnalysisType(kind),
		Result:     result,
	})
	if err != nil {
		return nil, err
	}

	return &CollectionAnalysis{
		ID:        row.ID,
		Type:      kind,
		Result:    row.Result,
		CreatedAt: row.CreatedAt,
	}, nil
}

func (core Core) ensureExtractions(
	ctx context.Context,
	userID uuid.UUID,
	collectionID uuid.UUID,
) error {

	q := core.Queries

	docs, err := q.GetCollectionDocuments(ctx, sqlgen.GetCollectionDocumentsParams{
		UserID:       userID,
		CollectionID: collectionID,
	})
	if err != nil {
		return err
	}

	for _, doc := range docs {
		exists, err := q.HasDocumentExtraction(ctx, doc.ID)
		if err != nil {
			return err
		}
		if exists {
			continue
		}

		extraction, err := core.extractDocumentContent(ctx, doc)
		if err != nil {
			return err
		}

		if _, err := q.CreateDocumentExtraction(ctx, sqlgen.CreateDocumentExtractionParams{
			DocumentID: doc.ID,
			Content:    extraction.Content,
		}); err != nil {
			return err
		}
	}

	return nil
}

// DocumentTextExtraction represents the structured output from OCR/text extraction
type DocumentTextExtraction struct {
	Content string `json:"content"`
}

func (DocumentTextExtraction) Describe() string {
	return `{
	"content": "string - The complete text content extracted from the document, preserving structure and ordering"
}`
}

func (core Core) extractDocumentContent(
	ctx context.Context,
	doc sqlgen.Document,
) (*DocumentTextExtraction, error) {

	url, err := core.Services.Minio.PresignedGetObject(
		ctx,
		core.UploadBucket,
		doc.S3Location,
		core.PresignedExpiry,
		nil,
	)
	if err != nil {
		return nil, err
	}

	img, err := imageanalysis.NewImageAnalyzer[DocumentTextExtraction](imageanalysis.NewImageAnalyzerParams{
		AI: core.Services.OpenAI,
	})
	if err != nil {
		return nil, err
	}

	return img.ExtractText(ctx, url)
}

type CollectionSnapshot struct {
	ID              uuid.UUID
	CombinedContent string
}

func (core Core) createSnapshot(
	ctx context.Context,
	collectionID uuid.UUID,
) (*CollectionSnapshot, error) {

	q := core.Queries

	extractions, err := q.GetDocumentExtractionsByCollection(ctx, collectionID)
	if err != nil {
		return nil, err
	}

	var combined strings.Builder
	for _, e := range extractions {
		combined.WriteString(e)
		combined.WriteString("\n\n")
	}

	row, err := q.CreateCollectionSnapshot(ctx, sqlgen.CreateCollectionSnapshotParams{
		CollectionID:    collectionID,
		CombinedContent: combined.String(),
	})
	if err != nil {
		return nil, err
	}

	return &CollectionSnapshot{
		ID:              row.ID,
		CombinedContent: combined.String(),
	}, nil
}

func (core Core) runAnalysis(
	ctx context.Context,
	content string,
	kind sqlgen.AnalysisType,
) (json.RawMessage, error) {

	schema := SchemaForAnalysis(kind)
	instructions := analysisInstructions(kind, schema)

	resp, err := core.Services.OpenAI.Responses.New(ctx, responses.ResponseNewParams{
		Model:        ChatModel,
		Instructions: openai.String(instructions),
		Input: responses.ResponseNewParamsInputUnion{
			OfString: openai.String(content),
		},
	})
	if err != nil {
		return nil, err
	}

	return json.RawMessage(resp.OutputText()), nil
}

func (core Core) GetCollectionAnalyses(
	ctx context.Context,
	userID uuid.UUID,
	collectionID uuid.UUID,
) ([]CollectionAnalysis, error) {

	q := core.Queries

	if _, err := q.GetCollection(ctx, sqlgen.GetCollectionParams{
		UserID: userID,
		ID:     collectionID,
	}); err != nil {
		return nil, err
	}

	rows, err := q.GetCollectionAnalysesByCollection(ctx, collectionID)
	if err != nil {
		return nil, err
	}

	results := make([]CollectionAnalysis, 0, len(rows))
	for _, r := range rows {
		results = append(results, CollectionAnalysis{
			ID:        r.ID,
			Type:      r.Type,
			Result:    r.Result,
			CreatedAt: r.CreatedAt,
		})
	}

	return results, nil
}

func analysisInstructions(kind sqlgen.AnalysisType, schema string) string {
	return fmt.Sprintf(`
You are analyzing study materials from a university course.

Your task: %s

Rules:
- Base your response ONLY on the provided content
- Do not invent topics
- Follow the JSON schema EXACTLY
- Do not include markdown or newlines

Schema:
%s
`, taskDescription(kind), schema)
}

func taskDescription(kind sqlgen.AnalysisType) string {
	switch kind {

	case sqlgen.AnalysisTypeSummary:
		return `
Generate a concise summary of the material.
Focus on the main ideas, definitions, and themes.
Assume the reader is a student reviewing before class or an exam.
Keep it brief but accurate.
`

	case sqlgen.AnalysisTypeFlashcards:
		return `
Generate a set of study flashcards.
Each flashcard should test a single concept, definition, or fact.
Questions should be clear and unambiguous.
Answers should be short, correct, and directly supported by the material.
`

	case sqlgen.AnalysisTypeQuiz:
		return `
Generate a short quiz to test understanding of the material.
Questions should cover a range of difficulties.
Prefer conceptual understanding over memorization.
Do not include trick questions.
`

	case sqlgen.AnalysisTypeDeepSummary:
		return `
Generate a comprehensive, in-depth explanation of the material.
Break the content into distinct concepts or topics.
Explain each concept clearly as if teaching it to a student encountering it for the first time.
Include definitions, explanations, and relationships between ideas where appropriate.
`

	default:
		return `
Analyze the provided material and extract its key educational content.
`
	}
}

package core

import (
	"context"
	"encoding/json"
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

// INTERNAL

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

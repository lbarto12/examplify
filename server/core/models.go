package core

import (
	"encoding/json"
	"server/sqlc/sqlgen"
	"time"

	"github.com/google/uuid"
)

type Collection struct {
	ID     uuid.UUID
	Title  string
	Course string
	Type   string
}

type Document struct {
	ID           uuid.UUID
	CollectionID uuid.UUID
	Title        string
	MimeType     string
	S3Location   string
}

type AnalysisType string

const (
	AnalysisSummary     AnalysisType = "summary"
	AnalysisFlashcards  AnalysisType = "flashcards"
	AnalysisQuiz        AnalysisType = "quiz"
	AnalysisDeepSummary AnalysisType = "deep_summary"
)

type CollectionAnalysis struct {
	ID        uuid.UUID           `json:"id"`
	Type      sqlgen.AnalysisType `json:"type"`
	Result    json.RawMessage     `json:"result"`
	CreatedAt time.Time           `json:"createdAt"`
}

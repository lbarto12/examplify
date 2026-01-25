package core

import "server/sqlc/sqlgen"

// SchemaForAnalysis returns a JSON schema string appropriate for the given AnalysisType.
// The schema is used to instruct the AI model how to structure its response.
func SchemaForAnalysis(kind sqlgen.AnalysisType) string {
	switch kind {
	case sqlgen.AnalysisTypeSummary:
		return `{
  "type": "object",
  "properties": {
    "summary": { "type": "string" }
  },
  "required": ["summary"]
}`

	case sqlgen.AnalysisTypeFlashcards:
		return `{
  "type": "array",
  "items": {
    "type": "object",
    "properties": {
      "question": { "type": "string" },
      "answer": { "type": "string" }
    },
    "required": ["question", "answer"]
  }
}`

	case sqlgen.AnalysisTypeQuiz:
		return `{
  "type": "array",
  "items": {
    "type": "object",
    "properties": {
      "question": { "type": "string" },
      "options": { "type": "array", "items": { "type": "string" } },
      "correct_index": { "type": "integer" }
    },
    "required": ["question", "options", "correct_index"]
  }
}`

	case sqlgen.AnalysisTypeDeepSummary:
		return `{
  "type": "array",
  "items": {
    "type": "object",
    "properties": {
      "concept": { "type": "string" },
      "definition": { "type": "string" },
      "details": { "type": "string" }
    },
    "required": ["concept", "definition", "details"]
  }
}`

	default:
		// Fallback: simple string
		return `{
  "type": "object",
  "properties": {
    "text": { "type": "string" }
  },
  "required": ["text"]
}`
	}
}

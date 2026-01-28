package core

import (
	"fmt"
	"server/sqlc/sqlgen"
)

// analysisInstructions generates AI instructions for a given analysis type
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

// taskDescription returns the task description for a given analysis type
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

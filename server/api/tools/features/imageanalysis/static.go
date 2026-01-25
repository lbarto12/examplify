package imageanalysis

import "fmt"

func aiChatbotAnalysisInstructions(schema string) string {
	return fmt.Sprintf(`
		you are a chatbot responsible for interpreting a provided image and expressing its pertinent data in JSON format.

		you are to use this EXACT schema. In your response do not use newline characters to prettify the output. state it simply on
		one line:

		Your schema:

		%s
	`,
		schema,
	)
}

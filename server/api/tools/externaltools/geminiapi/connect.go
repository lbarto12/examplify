package geminiapi

import (
	"context"
	"server/environment"

	"google.golang.org/genai"
)

func Connect(env *environment.Vars) (*genai.Client, error) {
	client, err := genai.NewClient(context.Background(), &genai.ClientConfig{
		APIKey: env.GeminiAIKey,
	})
	if err != nil {
		return nil, err
	}

	return client, nil
}

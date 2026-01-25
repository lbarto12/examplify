package gptapi

import (
	"server/environment"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
)

func Connect(env *environment.Vars) (*openai.Client, error) {
	client := openai.NewClient(option.WithAPIKey(env.OpenAIKey))
	return &client, nil
}

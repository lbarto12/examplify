package imageanalysis

import (
	"context"
	"errors"
	"net/url"

	"github.com/openai/openai-go/v3"
)

type describable_type interface {
	Describe() string
}

type image_analyzer_interface[T describable_type] interface {
	AnalyzeURL(ctx context.Context, kind AnalysisType, imageURL *url.URL) (*T, error)

	// Internal
	queryImageURL(ctx context.Context, request AIQueryImageParams) (*string, error)
}

type ImageAnalyzer[T describable_type] struct {
	AI *openai.Client
}

type NewImageAnalyzerParams struct {
	AI *openai.Client
}

func NewImageAnalyzer[T describable_type](data NewImageAnalyzerParams) (*ImageAnalyzer[T], error) {

	// Validate against interface
	var ftr image_analyzer_interface[T] = &ImageAnalyzer[T]{
		AI: data.AI,
	}
	return ftr.(*ImageAnalyzer[T]), nil
}

const ( // Consts
	ChatModel = openai.ChatModelGPT5ChatLatest
)

var ( // Errors
	ErrNoSchema error = errors.New("no schema specified")
)

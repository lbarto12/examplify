package imageanalysis

import (
	"context"
	"database/sql"
	"errors"
	"net/url"

	"github.com/minio/minio-go/v7"
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
	ObjectStore *minio.Client
	AI          *openai.Client
	Postgres    *sql.DB
}

type NewImageAnalyzerParams struct {
	ObjectStore *minio.Client
	AI          *openai.Client
	Postgres    *sql.DB
}

func NewImageAnalyzer[T describable_type](data NewImageAnalyzerParams) (*ImageAnalyzer[T], error) {

	// Validate against interface
	var ftr image_analyzer_interface[T] = &ImageAnalyzer[T]{
		ObjectStore: data.ObjectStore,
		AI:          data.AI,
		Postgres:    data.Postgres,
	}
	return ftr.(*ImageAnalyzer[T]), nil
}

const ( // Consts
	ChatModel = openai.ChatModelGPT5ChatLatest
)

var ( // Errors
	ErrNoSchema error = errors.New("no schema specified")
)

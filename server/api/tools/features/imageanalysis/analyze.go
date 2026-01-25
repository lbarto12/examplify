package imageanalysis

import (
	"context"
	"encoding/json"
	"net/url"
)

func (ftr ImageAnalyzer[T]) AnalyzeURL(ctx context.Context, imageURL *url.URL) (*T, error) {
	var x T

	response, err := ftr.queryImageURL(ctx, AIQueryImageParams{
		Instructions: aiChatbotAnalysisInstructions(x.Describe()),
		ImageURL:     imageURL,
	})
	if err != nil {
		return nil, err
	}

	var result T
	if err := json.Unmarshal([]byte(*response), &result); err != nil {
		return nil, err
	}

	return &result, nil
}

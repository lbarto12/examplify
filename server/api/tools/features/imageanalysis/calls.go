package imageanalysis

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/responses"
)

type AIQueryImageParams struct {
	Instructions string
	ImageURL     *url.URL
}

func (ftr ImageAnalyzer[T]) queryImageURL(ctx context.Context, request AIQueryImageParams) (*string, error) {
	if request.ImageURL == nil {
		return nil, fmt.Errorf("no ImageURL provided")
	}

	// 1️⃣ Download image
	httpResp, err := http.Get(request.ImageURL.String())
	if err != nil {
		return nil, fmt.Errorf("failed to download image: %w", err)
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to download image: status %d", httpResp.StatusCode)
	}

	imageBytes, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read image bytes: %w", err)
	}

	// 2️⃣ Encode as base64 data URL
	mimeType := httpResp.Header.Get("Content-Type")
	if mimeType == "" {
		mimeType = "application/octet-stream"
	}
	dataURL := fmt.Sprintf("data:%s;base64,%s", mimeType, base64.StdEncoding.EncodeToString(imageBytes))

	// 3️⃣ Prepare input for OpenAI
	inputImageParam := &responses.ResponseInputImageParam{
		Type:     "input_image",                          // required
		Detail:   responses.ResponseInputImageDetailAuto, // choose appropriate detail level
		ImageURL: openai.String(dataURL),                 // send bytes as data URL
	}

	// 4️⃣ Send request
	response, err := ftr.AI.Responses.New(ctx, responses.ResponseNewParams{
		Model:        ChatModel,
		Instructions: openai.String(request.Instructions),
		Input: responses.ResponseNewParamsInputUnion{
			OfInputItemList: responses.ResponseInputParam{
				responses.ResponseInputItemUnionParam{
					OfMessage: &responses.EasyInputMessageParam{
						Role: "user",
						Type: "message",
						Content: responses.EasyInputMessageContentUnionParam{
							OfInputItemContentList: responses.ResponseInputMessageContentListParam{
								{
									OfInputImage: inputImageParam,
								},
							},
						},
					},
				},
			},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("OpenAI request failed: %w", err)
	}

	result := response.OutputText()
	return &result, nil
}

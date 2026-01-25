package imageanalysis

import (
	"context"
	"log"
	"net/url"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/responses"
)

type AIQueryImageParams struct {
	Instructions string
	ImageURL     *url.URL
}

func (ftr ImageAnalyzer[T]) queryImageURL(ctx context.Context, request AIQueryImageParams) (*string, error) {

	response, err := ftr.AI.Responses.New(ctx, responses.ResponseNewParams{
		Model:        ChatModel,
		Instructions: openai.String(request.Instructions),
		Input: responses.ResponseNewParamsInputUnion{
			OfInputItemList: responses.ResponseInputParam{
				responses.ResponseInputItemUnionParam{
					OfMessage: &responses.EasyInputMessageParam{
						Role: "user",
						Type: "message", // must always be "message"
						Content: responses.EasyInputMessageContentUnionParam{
							OfInputItemContentList: responses.ResponseInputMessageContentListParam{
								{
									OfInputImage: &responses.ResponseInputImageParam{
										ImageURL: openai.String("https://docelf.com/images/hourly_invoice_template.png"),
									},
								},
							},
						},
					},
				},
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	ftr.AI.Responses.New(ctx, responses.ResponseNewParams{
		Model: ChatModel,
		Input: responses.ResponseNewParamsInputUnion{
			OfInputItemList: responses.ResponseInputParam{},
		},
	})
	if err != nil {
		return nil, err
	}

	result := response.OutputText()

	return &result, nil
}

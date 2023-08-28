package pkg

import (
	"context"

	"cloud.google.com/go/dialogflow/apiv2/dialogflowpb"
)

type dialogflowWrapperInterface interface {
	IndentDetectText(ctx context.Context, text, userSessionID string) (*dialogflowpb.QueryResult, error)
}

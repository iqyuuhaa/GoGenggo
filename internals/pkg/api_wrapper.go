package pkg

import (
	"context"
)

type apiWrapperInterface interface {
	Get(ctx context.Context, fullURI string, headers map[string]interface{}, params map[string]interface{}) ([]byte, error)
	Post(ctx context.Context, fullURI string, headers map[string]interface{}, params map[string]interface{}) ([]byte, error)
}

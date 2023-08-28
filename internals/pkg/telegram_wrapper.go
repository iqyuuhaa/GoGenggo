package pkg

import (
	"context"
)

type telegramWrapperInterface interface {
	SendMessage(ctx context.Context, chatFromID int64, message string) error
	GetUpdates(ctx context.Context, offset, limit int64, timeout int, allowUpdates []string) ([]byte, error)
}

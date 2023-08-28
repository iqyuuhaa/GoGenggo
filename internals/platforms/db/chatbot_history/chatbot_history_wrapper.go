package chatbot_history

import (
	"context"

	"gogenggo/internals/types/process"
)

type ChatbotHistoryWrapperInterface interface {
	InsertChatbotHistory(ctx context.Context, chatbotHistory process.ChatbotHistoryProcess) error
	InsertBulkChatbotHistory(ctx context.Context, chatbotHistories []process.ChatbotHistoryProcess) error
}

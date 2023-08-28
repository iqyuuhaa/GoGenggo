package chatbot_communication

import "context"

type ChatbotCommunicationWrapperInterface interface {
	GetResponse(ctx context.Context, chatKey string) (string, error)
}

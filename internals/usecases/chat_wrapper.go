package usecases

import (
	"context"

	"gogenggo/internals/types/payloads"
	"gogenggo/internals/types/responses"
)

type chatWrapperInterface interface {
	WebhookDialogflowUsecase(ctx context.Context, payload payloads.TelegramWebhookPayload) error
	TestDialogflowUsecase(ctx context.Context, payload payloads.DialogflowWebhookPayload) (responses.DialogflowWebhookResponse, error)

	// LiveChatBotUsecase(ctx context.Context, payload payloads.ChatBotPayload) error
}

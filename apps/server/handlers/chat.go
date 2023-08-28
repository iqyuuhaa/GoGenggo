package handlers

import (
	"log"
	"net/http"

	"gogenggo/apps/server/httputils"
	"gogenggo/internals/types/payloads"
)

func (h *Handler) HandleWebhookDialogflow(w http.ResponseWriter, r *http.Request) {
	var err error
	var payload payloads.TelegramWebhookPayload

	if err = httputils.GetFromBodyJSON(r, &payload); err != nil {
		log.Println("[Handlers - HandleWebhookDialogflow] Error Unmarshal request, err:", err)
		return
	}

	if err = h.usecases.Chat.WebhookDialogflowUsecase(r.Context(), payload); err != nil {
		log.Println("[Handlers - HandleWebhookDialogflow] Error processing webhook dialogflow usecase, err:", err)
	}
}

func (h *Handler) HandleTestDialogflow(w http.ResponseWriter, r *http.Request) {
	var err error
	var payload payloads.DialogflowWebhookPayload

	payload.Text, err = httputils.GetFromFormValue(r, "text", true)
	if err != nil {
		h.writeErrorResponse(w, r, err)
		return
	}

	response, err := h.usecases.Chat.TestDialogflowUsecase(r.Context(), payload)
	if err != nil {
		h.writeErrorResponse(w, r, err)
		return
	}

	h.writeSuccessResponse(w, r, response)
}

// func (h *Handler) HandleLiveChatBot(w http.ResponseWriter, r *http.Request) {
// var payload payloads.ChatBotPayload

// if err := httputils.GetFromBodyJSON(r, &payload); err != nil {
// 	log.Println("[Handlers - HandleLiveChatBot] Error Unmarshal request, err:", err)
// 	return
// }

// if payload.Type != constants.MessageType {
// 	return
// }

// if err := h.usecases.Chat.LiveChatBotUsecase(r.Context(), payload); err != nil {
// 	log.Println("[Handlers - HandleLiveChatBot] Error running live chat bot usecase, err:", err)
// }
// }

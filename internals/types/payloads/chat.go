package payloads

// Main payload
type TelegramWebhookPayload struct {
	UpdateID int64                         `json:"update_id"`
	Message  TelegramWebhookMessagePayload `json:"message"`
}

type DialogflowWebhookPayload struct {
	Text string `json:"text"`
}

// Additional struct
type TelegramWebhookMessagePayload struct {
	MessageID int64                                   `json:"message_id"`
	From      TelegramWebhookMessageFromPayload       `json:"from"`
	Chat      TelegramWebhookMessageChatPayload       `json:"chat"`
	Date      int64                                   `json:"date"`
	Text      string                                  `json:"text"`
	Entities  []TelegramWebhookMessageEntitiesPayload `json:"entities"`
}

type TelegramWebhookMessageFromPayload struct {
	ID           int64  `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	LanguageCode string `json:"language_code"`
}

type TelegramWebhookMessageChatPayload struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	Type      string `json:"type"`
}

type TelegramWebhookMessageEntitiesPayload struct {
	Offset int    `json:"offset"`
	Length int    `json:"length"`
	Type   string `json:"type"`
}

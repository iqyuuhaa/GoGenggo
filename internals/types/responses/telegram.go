package responses

type LongPollingTelegramResponse struct {
	Ok     bool                                `json:"ok"`
	Result []LongPollingTelegramResultResponse `json:"result"`
}

type LongPollingTelegramResultResponse struct {
	UpdateID int64                              `json:"update_id"`
	Message  LongPollingTelegramMessageResponse `json:"message"`
}

type LongPollingTelegramMessageResponse struct {
	MessageID int64                           `json:"message_id"`
	From      LongPollingTelegramFromResponse `json:"from"`
	Chat      LongPollingTelegramChatResponse `json:"chat"`
	Date      int64                           `json:"date"`
	Text      string                          `json:"text"`
}

type LongPollingTelegramFromResponse struct {
	ID           int64  `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	LanguageCode string `json:"language_code"`
}

type LongPollingTelegramChatResponse struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	Type      string `json:"type"`
}

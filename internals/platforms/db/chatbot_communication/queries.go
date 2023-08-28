package chatbot_communication

const (
	getChatResponseQuery = `
		SELECT chat_response
		FROM chatbot_communication
		WHERE chat_key = $1
	`
)

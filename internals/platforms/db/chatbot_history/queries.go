package chatbot_history

const (
	insertChatbotHistoryQuery = `
		INSERT INTO chatbot_history (method, identifier, datetime, process_time, dialogflow_process_time, non_dialogflow_process_time, send_message_process_time)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	insertBulkChatbotHistoryQuery = `
		INSERT INTO chatbot_history (method, identifier, datetime, process_time, dialogflow_process_time, non_dialogflow_process_time, send_message_process_time)
		VALUES 
	`
)

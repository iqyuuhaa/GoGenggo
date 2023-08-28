package chatbot_communication

import "github.com/jmoiron/sqlx"

type chatbotCommunicationQueries struct {
	getChatResponse *sqlx.Stmt
}

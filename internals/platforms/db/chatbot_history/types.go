package chatbot_history

import "github.com/jmoiron/sqlx"

type chatbotHistoryQueries struct {
	insertChatbotHistory *sqlx.Stmt
}

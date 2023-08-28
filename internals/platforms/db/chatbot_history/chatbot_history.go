package chatbot_history

import (
	"context"
	"log"
	"strings"
	"time"

	"gogenggo/config"
	"gogenggo/internals/types/constants"
	"gogenggo/internals/types/process"
	"gogenggo/modules/database"
)

func InitChatbotHistoryQueries() *chatbotHistoryQueries {
	return &chatbotHistoryQueries{
		insertChatbotHistory: database.Prepare(insertChatbotHistoryQuery),
	}
}

func (dbQueries *chatbotHistoryQueries) InsertChatbotHistory(ctx context.Context, chatbotHistory process.ChatbotHistoryProcess) error {
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(config.Configs.DB.Setting.Timeout)*time.Second)
	defer cancel()

	if config.Configs.DB.Setting.IsMaintenance {
		log.Println("[DB - ChatbotHistory - InsertChatbotHistory] DB is maintenance mode")
		return constants.ErrorDBMaintenance
	}

	if _, err := dbQueries.insertChatbotHistory.ExecContext(ctxTimeout,
		chatbotHistory.Method,
		chatbotHistory.Identifier,
		chatbotHistory.Datetime,
		chatbotHistory.ProcessTime,
		chatbotHistory.DialogflowProcessTime,
		chatbotHistory.NonDialogflowProcessTime,
		chatbotHistory.SendMessageProcessTime,
	); err != nil {
		log.Printf("[DB - ChatbotHistory - InsertChatbotHistory] Error insert chatbot history, err: %v", err)
		return constants.ErrorInternalServer
	}

	return nil
}

func (dbQueries *chatbotHistoryQueries) InsertBulkChatbotHistory(ctx context.Context, chatbotHistories []process.ChatbotHistoryProcess) error {
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(config.Configs.DB.Setting.Timeout)*time.Second)
	defer cancel()

	if config.Configs.DB.Setting.IsMaintenance {
		log.Println("[DB - ChatbotHistory - InsertBulkChatbotHistory] DB is maintenance mode")
		return constants.ErrorDBMaintenance
	}

	tx, err := database.DB.Begin()
	if err != nil {
		log.Println("[DB - ChatbotHistory - InsertBulkChatbotHistory] Error begining transaction, error:", err)
		return constants.ErrorInternalServer
	}

	queryStrings, queryArgs := formatingBulkInsertChatbotHistory(chatbotHistories)
	mainQuery := insertBulkChatbotHistoryQuery + strings.Join(queryStrings, ",")
	preparedQuery, err := tx.Prepare(mainQuery)
	if err != nil {
		tx.Rollback()
		log.Println("[DB - ChatbotHistory - InsertBulkChatbotHistory] Error preparing bulk insert chatbot history query, err:", err)
		return constants.ErrorInternalServer
	}

	if _, err = preparedQuery.ExecContext(ctxTimeout, queryArgs...); err != nil {
		tx.Rollback()
		log.Println("[DB - ChatbotHistory - InsertBulkChatbotHistory] Error running bulk insert chatbot history query, err:", err)
		return constants.ErrorInternalServer
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		log.Println("[DB - ChatbotHistory - InsertBulkChatbotHistory] Error commiting transaction, error:", err)
		return constants.ErrorInternalServer
	}

	return nil
}

func formatingBulkInsertChatbotHistory(chatbotHistories []process.ChatbotHistoryProcess) ([]string, []interface{}) {
	queryStrings := []string{}
	queryArgs := []interface{}{}

	for _, ch := range chatbotHistories {
		queryStrings = append(queryStrings, "(?, ?, ?, ?, ?, ?, ?)")
		queryArgs = append(queryArgs,
			ch.Method,
			ch.Identifier,
			ch.Datetime,
			ch.ProcessTime,
			ch.DialogflowProcessTime,
			ch.NonDialogflowProcessTime,
			ch.SendMessageProcessTime,
		)
	}

	return queryStrings, queryArgs
}

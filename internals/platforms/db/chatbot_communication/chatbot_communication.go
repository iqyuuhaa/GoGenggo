package chatbot_communication

import (
	"context"
	"log"
	"time"

	"gogenggo/config"
	"gogenggo/internals/types/constants"
	"gogenggo/modules/database"
)

func InitChatbotCommunicationQueries() *chatbotCommunicationQueries {
	return &chatbotCommunicationQueries{
		getChatResponse: database.Prepare(getChatResponseQuery),
	}
}

func (dbQueries *chatbotCommunicationQueries) GetResponse(ctx context.Context, chatKey string) (result string, err error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(config.Configs.DB.Setting.Timeout)*time.Second)
	defer cancel()

	if config.Configs.DB.Setting.IsMaintenance {
		log.Println("[DB - ChatbotCommunication - GetResponse] DB is maintenance mode")
		return result, constants.ErrorDBMaintenance
	}

	if err := dbQueries.getChatResponse.GetContext(ctxTimeout, &result, chatKey); err != nil {
		log.Println("[DB - ChatbotCommunication - GetResponse] Error getting chat response, err:", err)
		return result, constants.ErrorInternalServer
	}

	return result, nil
}

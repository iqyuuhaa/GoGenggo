package db

import (
	chatbotCommunication "gogenggo/internals/platforms/db/chatbot_communication"
	chatbotHistory "gogenggo/internals/platforms/db/chatbot_history"
)

var dbObjects *DBWrapper

func Init() *DBWrapper {
	if dbObjects != nil {
		return dbObjects
	}

	dbWrapper := new(DBWrapper)
	dbWrapper.ChatbotCommunication = chatbotCommunication.InitChatbotCommunicationQueries()
	dbWrapper.ChatbotHistory = chatbotHistory.InitChatbotHistoryQueries()

	dbObjects = dbWrapper

	return dbObjects
}

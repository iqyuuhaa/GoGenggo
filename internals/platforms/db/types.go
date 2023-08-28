package db

import (
	chatbotCommunication "gogenggo/internals/platforms/db/chatbot_communication"
	chatbotHistory "gogenggo/internals/platforms/db/chatbot_history"
)

type DBWrapper struct {
	ChatbotCommunication chatbotCommunication.ChatbotCommunicationWrapperInterface
	ChatbotHistory       chatbotHistory.ChatbotHistoryWrapperInterface
}

package utils

import (
	"os"

	"gogenggo/internals/types/constants"
)

func GetEnvironment() string {
	return os.Getenv(constants.GolangChatbotEnvironment)
}

func GetEnv(key string) string {
	return os.Getenv(key)
}

func SetEnv(key, value string) error {
	return os.Setenv(key, value)
}

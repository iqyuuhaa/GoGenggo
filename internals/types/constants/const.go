package constants

type ctxKey string

const AppsAsciiLogo = `
 ____  _  _     _   _                       __     ___ _ _                    ____                               _       
|___ \| || |   | | | | ___  _   _ _ __ ___  \ \   / (_) | | __ _  __ _  ___  / ___| _   _ _ __  _ __   ___  _ __| |_ ___ 
  __) | || |_  | |_| |/ _ \| | | | '__/ __|  \ \ / /| | | |/ _' |/ _' |/ _ \ \___ \| | | | '_ \| '_ \ / _ \| '__| __/ __|
 / __/|__   _| |  _  | (_) | |_| | |  \__ \   \ V / | | | | (_| | (_| |  __/  ___) | |_| | |_) | |_) | (_) | |  | |_\__ \
|_____|  |_|   |_| |_|\___/ \__,_|_|  |___/    \_/  |_|_|_|\__,_|\__, |\___| |____/ \__,_| .__/| .__/ \___/|_|   \__|___/
                                                                 |___/                   |_|   |_|                       
`

// Letter Runes
const LetterRunes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

// Apps Environments
const (
	DevelopmentEnvironment = "development"
	ProductionEnvironment  = "production"
)

// Environment Keys
const (
	GolangChatbotEnvironment     = "GOLANG_CHATBOT_ENV"
	GolangChatbotSecret          = "GOLANG_CHATBOT_SECRET"
	GolangChatbotSecretKey       = "GOLANG_CHATBOT_SECRET_KEY"
	GoogleApplicationCredentials = "GOOGLE_APPLICATION_CREDENTIALS"

	APIConfigPath = "API_CONFIG_PATH"
)

// Database
const (
	DatabaseType = "mysql"
)

// Context Key
const (
	StartTimeCtx = ctxKey("ctx_start_time")
)

// Language Type
const (
	LangID = "id"
	LangEN = "en"
)

const (
	WebhookMethod     = "webhook"
	LongPollingMethod = "long-polling"
)

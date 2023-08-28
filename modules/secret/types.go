package secret

type SecretObject struct {
	DB         DBSecretObject         `json:"database"`
	Dialogflow DialogflowSecretObject `json:"dialogflow"`
	Telegram   TelegramSecretObject   `json:"telegram"`
}

type DBSecretObject struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"db_name"`
	SSLMode  string `json:"ssl_mode"`
}

type DialogflowSecretObject struct {
	ProjectID string `json:"project_id"`
}

type TelegramSecretObject struct {
	Host     string `json:"host"`
	APIToken string `json:"api_token"`
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
}

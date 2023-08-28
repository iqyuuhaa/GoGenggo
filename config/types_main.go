package config

type SystemConfig struct {
	IsMaintenance         bool `ini:"IsMaintenance"`
	IsActiveWebhook       bool `ini:"IsActiveWebhook"`
	IsUsingDialogflow     bool `ini:"IsUsingDialogflow"`
	LongPollingPeriodTime int  `ini:"LongPollingPeriodTime"`
}

type MainHttpConfig struct {
	BaseURL      string `ini:"BaseURL"`
	Port         int    `ini:"Port"`
	WriteTimeout int    `ini:"WriteTimeout"`
	ReadTimeout  int    `ini:"ReadTimeout"`
}

type CacheConfig struct {
	TotalAPIRateLimit    int `ini:"TotalAPIRateLimit"`
	DurationAPIRateLimit int `ini:"DurationAPIRateLimit"`
	LifetimeSessionID    int `ini:"LifetimeSessionID"`
}

type BannedConfig struct {
	MaxCounter int `ini:"MaxCounter"`
}

type MainGeneratorConfig struct {
	SaltLimit int `ini:"SaltLimit"`
}

package config

type DBSettingConfig struct {
	IsMaintenance bool `ini:"IsMaintenance"`
	Timeout       int  `ini:"Timeout"`
	MaxIdleTime   int  `ini:"MaxIdleTime"`
	MaxLifetime   int  `ini:"MaxLifetime"`
	MaxIdleConns  int  `ini:"MaxIdleConns"`
	MaxOpenConns  int  `ini:"MaxOpenConns"`
}

type DBConnectionConfig struct {
	Scheme string `ini:"Scheme"`
}

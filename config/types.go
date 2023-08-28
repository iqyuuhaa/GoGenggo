package config

type Config struct {
	Main    MainConfig
	DB      DBConfig
	Message MessageConfig
}

type MainConfig struct {
	Test      TestConfig
	System    SystemConfig
	Http      MainHttpConfig
	Cache     CacheConfig
	Banned    BannedConfig
	Generator MainGeneratorConfig
}

type DBConfig struct {
	Test       TestConfig
	Setting    DBSettingConfig
	Connection DBConnectionConfig
}

type MessageConfig struct {
	Test    TestConfig
	Default DefaultConfig
}

// For testing get value
type TestConfig struct {
	Value string `ini:"Value"`
}

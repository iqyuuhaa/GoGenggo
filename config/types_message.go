package config

type DefaultConfig struct {
	MaintenanceModeMessage string `ini:"MaintenanceModeMessage"`
	BannedMessage          string `ini:"BannedMessage"`
}

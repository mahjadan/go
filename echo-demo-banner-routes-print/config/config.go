package config

import "os"

const Environment = "ENV"

func GetEnvironment() string {
	return getStringConfig(Environment, "dev")
}

func getStringConfig(configKey string, defaultValue string) (config string) {
	config = os.Getenv(configKey)
	if config == "" {
		return defaultValue
	}
	return config
}

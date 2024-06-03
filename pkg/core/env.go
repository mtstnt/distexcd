package core

import "os"

// Configures values by Environment Vars.

var (
	ExposePort string
)

func ConfigureEnvironment() error {
	ExposePort = getEnvOrDefault("EXPOSE_PORT", "4999")
	return nil
}

func mustGetEnv(key string) string {
	value, isExist := os.LookupEnv(key)
	if !isExist {
		GetLogger().Fatal("Failed to get required field \"" + key + "\" in environment variables")
	}
	return value
}

func getEnvOrDefault(key, defaultValue string) string {
	value, isExist := os.LookupEnv(key)
	if !isExist {
		return defaultValue
	}
	return value
}

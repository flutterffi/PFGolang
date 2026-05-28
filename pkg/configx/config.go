package configx

import (
	"os"
	"strconv"
)

type AppConfig struct {
	AppName string
	Port    int
	Debug   bool
}

func Load() AppConfig {
	return AppConfig{
		AppName: getenv("PFGOLANG_APP_NAME", "PFGolang"),
		Port:    getenvInt("PFGOLANG_PORT", 8080),
		Debug:   getenvBool("PFGOLANG_DEBUG", false),
	}
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func getenvInt(key string, fallback int) int {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	parsed, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}

	return parsed
}

func getenvBool(key string, fallback bool) bool {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	parsed, err := strconv.ParseBool(value)
	if err != nil {
		return fallback
	}

	return parsed
}

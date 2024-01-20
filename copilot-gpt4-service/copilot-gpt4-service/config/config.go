package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port         string
	Cache        bool
	CachePath    string
	Host         string
	Debug        bool
	Logging      bool
	LogLevel     string
	CopilotToken string
}

var ConfigInstance *Config = NewConfig()

func NewConfig() *Config {
	// if exists config.env, load it
	if _, err := os.Stat("config.env"); err == nil {
		err := godotenv.Load("config.env")
		if err != nil {
			fmt.Println("Error loading config.env file")
		}
	}

	return &Config{
		Host:         getEnvOrDefault("HOST", "localhost"),
		Port:         getEnvOrDefault("PORT", "8080"),
		Cache:        getEnvOrDefaultBool("CACHE", true),
		CachePath:    getEnvOrDefault("CACHE_PATH", "db/cache.sqlite3"),
		Debug:        getEnvOrDefaultBool("DEBUG", false),
		Logging:      getEnvOrDefaultBool("LOGGING", true),
		LogLevel:     getEnvOrDefault("LOG_LEVEL", "info"),
		CopilotToken: getEnvOrDefault("COPILOT_TOKEN", ""),
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getEnvOrDefaultBool(key string, defaultValue bool) bool {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	s, err := strconv.ParseBool(value)
	if err != nil {
		return false
	}
	return s
}

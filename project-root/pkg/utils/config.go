package config

import (
	"os"
)

type Config struct {
	TelegramToken string
	DBPath        string
	APIAddr       string
	TLSCert       string
	TLSKey        string
	LogLevel      string
}

func LoadConfig() Config {
	return Config{
		TelegramToken: getEnv("TELEGRAM_TOKEN", "YOUR_TELEGRAM_TOKEN"),
		DBPath:        getEnv("DB_PATH", "project-root/database.db"),
		APIAddr:       getEnv("API_ADDR", ":8443"),
		TLSCert:       getEnv("TLS_CERT", "path/to/cert.pem"),
		TLSKey:        getEnv("TLS_KEY", "path/to/key.pem"),
		LogLevel:      getEnv("LOG_LEVEL", "INFO"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

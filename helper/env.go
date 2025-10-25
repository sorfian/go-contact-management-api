package helper

import (
	"os"
	"strconv"
	"time"
)

func GetEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func GetEnvAsInt(key string, defaultValue int) int {
	valueStr := GetEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}

func GetEnvAsDuration(key string, defaultValue time.Duration) time.Duration {
	valueStr := GetEnv(key, "")
	if value, err := time.ParseDuration(valueStr); err == nil {
		return value
	}
	return defaultValue
}

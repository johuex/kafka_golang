package config

import (
	"os"
	"strconv"
)

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getEnvInt(key string, defaultVal int) int {
	if value, exists := os.LookupEnv(key); exists {
		intVal, _ := strconv.Atoi(value)
		return intVal
	}

	return defaultVal
}

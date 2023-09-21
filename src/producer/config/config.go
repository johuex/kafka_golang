package config

import (
	"os"
	"strconv"
)

type Config struct {
	KafkaURL1  string
	KafkaURL2  string
	KafkaTopic string
	BatchSize  int
	TotalSize  int
}

func New() *Config {
	return &Config{
		KafkaURL1:  getEnv("KAFKA_URL_1", "localhost:29092"),
		KafkaURL2:  getEnv("KAFKA_URL_2", "localhost:39093"),
		KafkaTopic: getEnv("KAFKA_TOPIC", "example_topic"),
		BatchSize:  getEnvInt("BATCH_SIZE_PRODUCER", 100000),
		TotalSize:  getEnvInt("TOTAL_SIZE_PRODUCER", 10000000),
	}
}

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

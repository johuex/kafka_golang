package config

type Config struct {
	BrokerUrl   string
	BrokerPort  string
	BrokerTopic string
	GroupID     string
	BatchSize   int
}

func NewConfig() *Config {
	return &Config{
		BrokerUrl:   getEnv("KAFKA_URL", ""),
		BrokerPort:  getEnv("KAFKA_PORT", ""),
		BrokerTopic: getEnv("KAFKA_TOPIC", ""),
		GroupID:     getEnv("KAFKA_GROUP", ""),
		BatchSize:   getEnvInt("BATCH_SIZE", 100),
	}
}

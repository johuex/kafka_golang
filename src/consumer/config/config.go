package config

type Config struct {
	BrokerUrl     string
	BrokerPort    string
	BrokerTopic   string
	GroupID       string
	BatchSize     int
	MinBatchSize  int
	MaxBatchSize  int
	MaxReaderWait int
}

func NewConfig() *Config {
	return &Config{
		BrokerUrl:     getEnv("KAFKA_URL", ""),
		BrokerPort:    getEnv("KAFKA_PORT", ""),
		BrokerTopic:   getEnv("KAFKA_TOPIC", ""),
		GroupID:       getEnv("KAFKA_GROUP", ""),
		BatchSize:     getEnvInt("BATCH_SIZE_CONSUMER", 6250),
		MinBatchSize:  getEnvInt("KAFKA_MINBATCHSIZE", 168e4),
		MaxBatchSize:  getEnvInt("KAFKA_MAXBATCHSIZE", 200e4),
		MaxReaderWait: getEnvInt("KAFKA_MAXREADERWAIT", 10),
	}
}

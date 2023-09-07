package kafka_service

import (
	"context"
	"producer/config"
	"time"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/snappy"
)

type KafkaService struct {
	writer *kafka.Writer
	Config *config.Config
}

func (k *KafkaService) CreateConnection() {
	clientId := uuid.New().String()
	dialer := &kafka.Dialer{
		Timeout:  10 * time.Second,
		ClientID: clientId,
	}
	kafkaBrokerUrls := []string{k.Config.KafkaURL1, k.Config.KafkaURL2}
	config := kafka.WriterConfig{
		Brokers:          kafkaBrokerUrls,
		Topic:            k.Config.KafkaTopic,
		Balancer:         &kafka.CRC32Balancer{},
		Dialer:           dialer,
		WriteTimeout:     10 * time.Second,
		ReadTimeout:      10 * time.Second,
		CompressionCodec: snappy.NewCompressionCodec(),
		BatchSize:        k.Config.BatchSize / 4,
	}
	w := kafka.NewWriter(config)
	if w != nil {
		k.writer = w
	}
}

func (k *KafkaService) Push(message kafka.Message, ctx context.Context) error {
	if k.writer == nil {
		k.CreateConnection()
	}
	err := k.writer.WriteMessages(
		ctx,
		message,
	)
	return err
}

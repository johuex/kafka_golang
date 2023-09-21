package app_service

import (
	"context"
	"encoding/json"
	"log"
	"math/rand"
	"producer/config"
	"producer/models"
	"producer/services/kafka_service"
	"time"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

type Service struct {
	Config       *config.Config
	Logger       *log.Logger
	KafkaService *kafka_service.KafkaService
}

func (s *Service) Ping() []byte {
	res, err := json.Marshal(s.Config)
	if err != nil {
		s.Logger.Fatalf("Config encoding error: %v", err)
	}
	return res
}

func (s *Service) Kafka() {
	var i int
	for i = 0; i < s.Config.TotalSize/s.Config.BatchSize; i++ {
		messages := make([]kafka.Message, s.Config.BatchSize)
		for idx := range messages {
			messageKey := uuid.New().String()

			messageVal := models.Transaction{
				TransactionType: models.TransactionType(rand.Intn(7)),
				Location:        models.Locations[rand.Intn(5)],
				TransactionId:   rand.Int63(),
				AccountNumber:   rand.Int63(),
				Amount:          rand.Intn(110000),
				Time:            time.Now(),
			}
			jsonVal, _ := json.Marshal(messageVal)
			message := kafka.Message{
				Key:   []byte(messageKey),
				Value: jsonVal,
			}
			messages[idx] = message
		}

		err := s.KafkaService.Push(messages, context.Background())
		if err != nil {
			s.Logger.Fatal("failed to write messages:", err)
		}
	}

	s.Logger.Printf("Send to Kafka: %d batches with size %d", i, s.Config.BatchSize)
}

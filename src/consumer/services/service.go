package service

import (
	"context"
	"encoding/json"
	"log"
	"sync"
	"time"

	"consumer/config"
	"consumer/models"
	"consumer/repositories"

	"github.com/segmentio/kafka-go"
)

type Service struct {
	Logger                *log.Logger
	Config                *config.Config
	TransactionRepository *repositories.TransactionRepository
}

func (s *Service) Consume(wg *sync.WaitGroup) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{s.Config.BrokerUrl},
		Topic:       s.Config.BrokerTopic,
		GroupID:     s.Config.GroupID,
		Logger:      s.Logger,
		MinBytes:    s.Config.MinBatchSize,                               // min size of batch, each message size is 168 bytes
		MaxBytes:    s.Config.MaxBatchSize,                               // max size of batch
		MaxWait:     time.Duration(s.Config.MaxReaderWait) * time.Second, // max timeout to wait for read if not cross the min bytes
		StartOffset: kafka.FirstOffset,                                   // fetch like queue
	})
	ctx := context.Background()

	for {
		kafka_messages := make([]kafka.Message, s.Config.BatchSize)
		// read batch
		for idx := range kafka_messages {
			msg, err := r.FetchMessage(ctx) // block & wait for NewMessage
			if err != nil {
				s.Logger.Fatal("could not read message " + err.Error())
			}
			s.Logger.Printf("Received from master %s:%s ", string(msg.Key), string(msg.Value))
			kafka_messages[idx] = msg
		}

		//convert from bytes to orm
		orms := s.convertToOrm(kafka_messages)

		// write batch to repo
		err := s.TransactionRepository.InsertTransactions(orms)
		if err != nil {
			s.Logger.Fatal("error while insert to DB " + err.Error())
		}

		err = r.CommitMessages(ctx, kafka_messages...)
		if err != nil {
			s.Logger.Fatal("error while commiting kafka messages " + err.Error())
		}
	}
}

// TODO adapt the algorithm to MapReduce
//func mergeResults(firstArr []repositories.TransactionDB, secondArr []repositories.TransactionDB) {
//	firstArr = append(firstArr, secondArr...)
//}

func (s *Service) convertToOrm(messages []kafka.Message) []repositories.TransactionDB {
	db_models := make([]repositories.TransactionDB, len(messages))

	for idx, msg := range messages {
		transaction := models.Transaction{}
		json.Unmarshal(msg.Value, &transaction)
		accepted := true // default
		if transaction.Location == "LPP" || transaction.Amount > 100000 || transaction.TransactionType.String() == "Undefined" || transaction.TransactionType.String() == "RecurringPayment" {
			accepted = false
		}
		db_models[idx] = repositories.TransactionDB{
			TransactionType: transaction.TransactionType.String(),
			Location:        transaction.Location,
			TransactionId:   transaction.TransactionId,
			AccountNumber:   transaction.AccountNumber,
			Amount:          transaction.Amount,
			Time:            transaction.Time,
			Accepted:        accepted,
		}
	}
	return db_models
}

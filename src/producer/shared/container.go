package shared

import (
	"log"
	"os"
	"producer/config"
	"producer/services/app_service"
	"producer/services/kafka_service"
)

var ContainerItem Container

type Container struct {
	Service *app_service.Service
}

func init() {
	config_ := config.New()
	logger := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	kafkaService := kafka_service.KafkaService{Config: config_, Logger: logger}
	service := app_service.Service{
		Config: config_, Logger: logger, KafkaService: &kafkaService,
	}
	ContainerItem = Container{&service}
}

package shared

import (
	"consumer/config"
	"consumer/repositories"
	service "consumer/services"
	"log"
	"os"
)

var ContainerItem Container

type Container struct {
	Service *service.Service
}

func init() {
	logger := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	config_ := config.NewConfig()
	db_config := config.NewDBConfig()
	dbConn, _ := GetDBConnection(db_config)
	transaction_repo := repositories.NewTransactionRepository(dbConn)
	service_ := &service.Service{Logger: logger, Config: config_, TransactionRepository: transaction_repo}

	ContainerItem = Container{service_}
}

package shared

import (
	"consumer/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDBConnection(db_config *config.DBConfig) (*gorm.DB, error) {
	dsn := db_config.GetDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

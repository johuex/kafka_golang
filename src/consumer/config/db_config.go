package config

import "fmt"

type DBConfig struct {
	Host     string
	User     string
	DBName   string
	Password string
	SSLMode  string
	Port     int
}

func NewDBConfig() *DBConfig {
	return &DBConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		User:     getEnv("DB_USER", "dev_user"),
		DBName:   getEnv("DB_NAME", "dev_test"),
		Password: getEnv("DB_PASSWD", "123456"),
		Port:     getEnvInt("DB_PORT", 5432),
		SSLMode:  getEnv("DB_SSL", "disable"),
	}
}

func (db_c *DBConfig) GetDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		db_c.Host, db_c.User, db_c.Password, db_c.DBName, db_c.Port, db_c.SSLMode)
}

package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	DBName     string
	DBUsername string
	DBPassword string
	DBDatabase string
}
type ServiceConfig struct {
	Host string
	Port string
}

func LoadConfig() (DatabaseConfig, ServiceConfig) {
	DBConfig := initDBConfig()
	servConfig := initServiceConfig()
	return DBConfig, servConfig
}

func initDBConfig() DatabaseConfig {
	DBConfig := DatabaseConfig{
		DBName:     "hospital_management_system",
		DBUsername: "priyanshi.j",
		DBPassword: "my",
	}
	return DBConfig
}

func initServiceConfig() ServiceConfig {
	servConfig := ServiceConfig{
		Host: "localhost",
		Port: "5432",
	}
	return servConfig
}

func DatabaseConnection(DBConfig DatabaseConfig, servConfig ServiceConfig) *gorm.DB {
	sqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", servConfig.Host, servConfig.Port, DBConfig.DBUsername, DBConfig.DBPassword, DBConfig.DBName)
	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}
	return db
}

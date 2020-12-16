package db

import (
	"fmt"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DatabaseConfig is struct of config database
type DatabaseConfig struct {
	DBPort     string
	DBHost     string
	DBDatabase string
	DBUsername string
	DBPassword string
}

// Database is connection instance of database
type Database struct {
	Connection *gorm.DB
}

var once sync.Once
var database *Database

// Connect is method to get instance of database connection
func Connect() *Database {
	dsn := getDsn()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	if database == nil {
		once.Do(func() {
			database = &Database{
				Connection: db,
			}
		})
	}

	return database
}

func getDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		DBPort:     os.Getenv("DB_PORT"),
		DBDatabase: os.Getenv("DB_DATABASE"),
		DBHost:     os.Getenv("DB_HOST"),
		DBUsername: os.Getenv("DB_USERNAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),
	}
}

func getDsn() string {
	dbConfig := getDatabaseConfig()
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		dbConfig.DBHost,
		dbConfig.DBUsername,
		dbConfig.DBPassword,
		dbConfig.DBDatabase,
		dbConfig.DBPort,
	)
}

package config

import (
	"fmt"
	"log"
	"sync"

	"github.com/yet-another-todo-list-golang/model/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DatabaseConfig database minimal configuration
type DatabaseConfig struct {
	DBPort     string
	DBHost     string
	DBDatabase string
	DBUsername string
	DBPassword string
}

var connection *gorm.DB
var once sync.Once

// DatabaseConnect initialize database if not being initialize else get connection
func DatabaseConnect() *gorm.DB {
	if connection == nil {
		once.Do(func() {
			connection = initialize()
		})
	}

	return connection

}

// init initialize db function
func initialize() *gorm.DB {
	dsn := getDsn()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	migrate(db)
	log.Println("Database Connection is being initialized successfully")
	return db
}

// Migrate orm db wrapper for autoloading the table
func migrate(db *gorm.DB) {
	db.AutoMigrate(&entity.Todo{})
}

//getDatabaseConfig generate database config and return the config
func getDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		DBPort:     GetEnvConfig("db.port"),
		DBDatabase: GetEnvConfig("db.database"),
		DBHost:     GetEnvConfig("db.host"),
		DBUsername: GetEnvConfig("db.username"),
		DBPassword: GetEnvConfig("db.password"),
	}
}

//getDsn generate dsn and return the dsn
func getDsn() string {
	config := getDatabaseConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		config.DBHost,
		config.DBUsername,
		config.DBPassword,
		config.DBDatabase,
		config.DBPort,
	)
	return dsn
}

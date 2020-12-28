package db

import (
	"github.com/yet-another-todo-list-golang/config"
	"log"
	"sync"

	"github.com/yet-another-todo-list-golang/model/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var connection *gorm.DB
var once sync.Once

// DatabaseConnect initialize database if not being initialize else get connection
func GetConnection() *gorm.DB {
	if connection == nil {
		once.Do(func() {
			connection = initialize()
		})
	}

	return connection

}

// init initialize db function
func initialize() *gorm.DB {
	dsn := config.GetDsn()
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
	err := db.AutoMigrate(&entity.Todo{}, &entity.User{})
	if err != nil {
		log.Fatal(err.Error())
	}
}

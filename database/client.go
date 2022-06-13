package database

import (
	"golang-crud-rest-api/entities"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var connectionError error

func Connect(connectionString string) {
	Instance, connectionError = gorm.Open(
		postgres.Open(connectionString), &gorm.Config{})

	if connectionError != nil {
		log.Fatal(connectionError)
		panic("Can't connect to database")
	}

	log.Println("Connected to database")
}

func Migrate() {
	Instance.AutoMigrate(&entities.Inventory{})
	log.Println("Migration of database is completed")
}

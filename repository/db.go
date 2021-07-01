package repository

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB
var dbErr error

func ConnectToDatabase() {

	databaseLogin := "host=dock_db_postgres user=root password=password dbname=airways port=5432 sslmode=disable"

	Db, dbErr = gorm.Open(postgres.Open(databaseLogin), &gorm.Config{})
	if dbErr != nil {
		fmt.Printf("DB connection failed")
		log.Fatal(dbErr)
	}
}

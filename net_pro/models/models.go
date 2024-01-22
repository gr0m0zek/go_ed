package models

import (
	"fmt"

	"bd/bd.go"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func db_connect() {
	dsn := "host=localhost user=hugh password=SuperSecretPassword dbname=db port=5433 sslmode=disable"
	Db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err == nil {
		fmt.Printf("%T", Db)
	}
}

func db_migrate(db *gorm.DB) {

	migrator := db.Migrator()
	if !migrator.HasTable("vehicles") {
		migrator.CreateTable(&bd.Vehicle{})

	}
}

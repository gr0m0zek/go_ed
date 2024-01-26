package models

import (
	"fmt"
	"http/bd"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Db_connect() {
	dsn := fmt.Sprintf("host=%s user=hugh password=SuperSecretPassword dbname=db port=5432 sslmode=disable", os.Getenv("DB_HOST"))
	DbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	Db = DbConn
	if err == nil {
	}

	db_migrate(Db)
}

func db_migrate(db *gorm.DB) {
	migrator := db.Migrator()

	if !(migrator.HasTable("vehicles")) {
		migrator.CreateTable(&bd.Vehicle{})
	}
}

func Add_vehicle(v bd.Vehicle) {
	Db.Create(&v)
}

func Get_vehicle(id string) (bd.Vehicle, error) {
	var vehicle bd.Vehicle
	res := Db.First(&vehicle, "id = ?", id)
	err := res.Error
	if err != nil {
		return bd.Vehicle{}, err
	}
	return vehicle, err
}

func Get_all_vehicle() []bd.Vehicle {
	var vehicles []bd.Vehicle
	Db.Limit(-1).Find(&vehicles)
	return vehicles
}

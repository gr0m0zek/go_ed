package models

import (
	"context"

	"github.com/gin-gonic/gin"

	"log"
	"net_pro/bd"

	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var db *bun.DB

func Db_connect() {
	ctx := context.Background()

	dsn := "postgresql://hugh:SuperSecretPassword@localhost:5433/db?sslmode=disable"
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db = bun.NewDB(sqldb, pgdialect.New())

	if _, err := db.NewCreateTable().Model((*bd.Vehicle)(nil)).IfNotExists().Exec(ctx); err != nil {
		log.Println(err)
	}

	if err := db.Ping(); err != nil {
		log.Println(err)
		return
	}
}

func Get_vehicle(id string, context *gin.Context) (bd.Vehicle, error) {

	vehicle := new(bd.Vehicle)
	err := db.NewSelect().Model(vehicle).Where("id = ?", id).Scan(context)

	return *vehicle, err
}

func Get_all_vehicle(context *gin.Context) ([]bd.Vehicle, error) {

	vehicles := new([]bd.Vehicle)
	err := db.NewSelect().Model(vehicles).Scan(context)

	return *vehicles, err
}

func Add_vehicle(vehicle *bd.Vehicle, context *gin.Context) error {
	_, err := db.NewInsert().Model(vehicle).Exec(context)
	return err
}

func UpdateVehicle(vehicle *bd.Vehicle, context *gin.Context) error {
	_, err := db.NewUpdate().Model(vehicle).WherePK().Exec(context)
	return err
}

package main

import (
	"fmt"
	"net/http"

	"http/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var str = "before"
var Db *gorm.DB

func main() {
	db_connect()
	r := gin.Default()
	r.GET("/", getting)
	r.POST("/", postting)
	r.Run(":3000")

}

func getting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ping": str})
}

func postting(c *gin.Context) {
	// Db.(&,c.Query("id"))
	str += " " + c.Query("id")
}

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
		migrator.CreateTable(&models.Vehicle{})

	}
}

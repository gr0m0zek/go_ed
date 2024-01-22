package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var str = "before"
var Db *gorm.DB

func main() {
	// db_connect()
	// r := gin.Default()
	// r.GET("/", getting)
	// r.POST("/", postting)
	// r.Run(":3000")

}

func getting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ping": str})
}

func postting(c *gin.Context) {
	// Db.(&,c.Query("id"))
	str += " " + c.Query("id")
}

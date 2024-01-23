package main

import (
	"http/logic"
	"http/models"
)

// var Db *gorm.DB

func main() {
	models.Db_connect()
	logic.Run_back()
}

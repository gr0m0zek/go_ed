package main

import (
	"http/logic"
	"http/models"
)

func main() {
	models.Db_connect()
	logic.Run_back()
}

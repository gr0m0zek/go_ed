package main

import (
	"net_pro/logic"
	"net_pro/models"
)

func main() {
	models.Db_connect()
	logic.Run_back()
}

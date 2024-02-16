package handlers

import (
	"log"
	"net/http"
	"net_pro/bd"
	"net_pro/models"

	"github.com/gin-gonic/gin"
)

type UserRequest struct {
	Action  string     `json:"action"`
	Vehicle bd.Vehicle `json:"body"`
}

func Getting(context *gin.Context) {
	if context.Query("type") == "one" {

		id := context.Query("id")

		if id != "" {

			vehicle, err := models.Get_vehicle(id, context)

			if err != nil {
				context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			}

			context.JSON(http.StatusOK, vehicle)
		}
	}

	if context.Query("type") == "all" {

		vehicles, error := models.Get_all_vehicle(context)

		if error != nil {
			context.JSON(http.StatusNotFound, gin.H{"error": error.Error()})
			return
		}

		context.JSON(http.StatusOK, vehicles)
	}
}

func Posting(context *gin.Context) {

	var req UserRequest
	err := context.ShouldBind(&req)

	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var vehicle bd.Vehicle = req.Vehicle

	switch req.Action {

	case "add":
		err := models.Add_vehicle(&vehicle, context)

		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		context.JSON(http.StatusOK, gin.H{"message": "Vehicle created successfully", "vehicle": vehicle})

	case "update":
		err := models.UpdateVehicle(&vehicle, context)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusOK, gin.H{"message": "Vehicle updated successfully", "vehicle": vehicle})
	}

	log.Print(req, vehicle)
}

package handlers

import (
	"http/bd"
	"http/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserRequest struct {
	Action  string     `json:"action"`
	Vehicle bd.Vehicle `json:"body"`
}

func Getting(c *gin.Context) {
	if c.Query("type") == "one" {

		id := c.Query("id")

		if id != "" {
			vehicle, err := models.Get_vehicle(id)

			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, vehicle)
		}
	}

	if c.Query("type") == "all" {
		c.JSON(http.StatusOK, models.Get_all_vehicle())
	}
}

func Posting(c *gin.Context) {

	var req UserRequest
	err := c.ShouldBind(&req)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var vehicle bd.Vehicle = req.Vehicle

	switch req.Action {
	case "add":
		models.Add_vehicle(vehicle)
		c.JSON(200, gin.H{"message": "Vehicle created successfully", "vehicle": vehicle})

	case "update":
		err := models.UpdateVehicle(vehicle)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "Vehicle updated successfully", "vehicle": vehicle})
	}

	log.Print(req, vehicle)
}

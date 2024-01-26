package handlers

import (
	"http/models"
	"http/bd"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
	var vehicle bd.Vehicle

	if err := c.ShouldBind(&vehicle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.Add_vehicle(vehicle)
	c.JSON(http.StatusOK, gin.H{"message": "Vehicle created successfully", "vehicle" : vehicle})
}

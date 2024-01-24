package handlers

import (
	"http/models"
	"net/http"
	"strconv"

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

func Postting(c *gin.Context) {
	var dist uint64
	var year uint16

	if distance, err1 := strconv.Atoi(c.Query("distance")); err1 == nil {
		dist = uint64(distance)
	}

	if y, err2 := strconv.Atoi(c.Query("year")); err2 == nil {
		year = uint16(y)
	}

	models.Add_vehicle(c.Query("mark"), c.Query("model"), c.Query("number"), dist, year)
}

package handlers

import (
	"http/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Getting(c *gin.Context) {
	id := c.Query("id")

	if id != "" {
		c.JSON(http.StatusOK, models.Get_vehicle(id))
	}
}

func Postting(c *gin.Context) {
	if dist, err1 := strconv.Atoi(c.Query("distance")); err1 == nil {

		if year, err2 := strconv.Atoi(c.Query("year")); err2 == nil {
			models.Add_vehicle(c.Query("mark"), c.Query("model"), c.Query("number"), uint64(dist), uint16(year))
		}
	}
}

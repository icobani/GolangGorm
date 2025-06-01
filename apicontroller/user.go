/*
	UCR TECH
	User        : EUR
	Name        : Enes Emin Ucar
	Date        : 1.06.2025
	Time        : 17:57
	Notes       :


*/

package apicontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/icobani/GolangGorm/Services/ArrsPlanningService"
	"github.com/icobani/GolangGorm/config"
	"net/http"
)

func GetUsers(c *gin.Context) {
	search := c.Query("search")
	s := ArrsPlanningService.NewArrsPlanningService(config.Params, config.DB)
	users, err := s.GetUsers(search)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve users",
			"details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)

}

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

func PutChangeMyPassword(c *gin.Context) {
	var req ArrsPlanningService.ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data",
			"details": err.Error()})
		return
	}

	s := ArrsPlanningService.NewArrsPlanningService(config.Params, config.DB)
	if err := s.ChangeMyPassword(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to change password",
			"details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password changed successfully"})
}

/*
	UCR TECH
	User        : EUR
	Name        : Enes Emin Ucar
	Date        : 1.06.2025
	Time        : 17:54
	Notes       :


*/

package apiroots

import (
	"github.com/gin-gonic/gin"
	"github.com/icobani/GolangGorm/apicontroller"
)

func UserRoot(api *gin.RouterGroup) {

	api.GET("/users", apicontroller.GetUsers)

}

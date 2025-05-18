/*
	B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
	User        : ICI
	Name        : Ibrahim COBANI
	Date        : 18.05.2025
	Time        : 18:45
	Notes       : Taha


*/

package Test

import (
	"log"
	"testing"

	"github.com/icobani/GolangGorm/Services/ArrsPlanningService"
	"github.com/icobani/GolangGorm/config"
)

func TestArrsPlanning2Test(t *testing.T) {
	if err := config.InitConfigFile(); err != nil {
		log.Println("Error: ", err)

		return
	}

	if err := config.InitDB(); err != nil {
		log.Println("Error: ", err)
		return
	}

	service := ArrsPlanningService.NewArrsPlanningService(config.Params, config.DB)

	var req = ArrsPlanningService.RegisterUserRequest{
		FullName: "Ibrahim ÇOBANİ",
		UserName: "icobani",
		Password: "112233",
		Email:    "ibrahim@imaconsult.com",
		Phone:    "5325401194",
	}

	if user, err := service.RegisterUser(req); err != nil {
		log.Println("Error: ", err)
	} else {
		log.Println("User Created:")
		service.PrintPrettyStruct(user)
	}
}

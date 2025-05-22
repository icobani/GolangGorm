/*
	B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
	User        : ICI
	Name        : Ibrahim COBANI
	Date        : 18.05.2025
	Time        : 18:45
	Notes       : Enes


*/

package Test

import (
	"log"
	"testing"

	"github.com/icobani/GolangGorm/Services/ArrsPlanningService"
	"github.com/icobani/GolangGorm/config"
)

func TestArrsPlanningTest(t *testing.T) {
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

	req = ArrsPlanningService.RegisterUserRequest{
		FullName: "Enes Emin Uçar",
		UserName: "ensemio",
		Password: "070706",
		Email:    "ucarenesemin@gmail.com",
		Phone:    "5317887577",
	}

	user, err := service.RegisterUser(req)
	if err != nil {
		log.Println("Error: ", err)
	} else {
		log.Println("User Created:")
		service.PrintPrettyStruct(user)

	}

}

func TestChangeMyPassword(t *testing.T) {

	if err := config.InitConfigFile(); err != nil {
		t.Fatalf("Config yükelenemedi : %v", err)

		return

	}

	if err := config.InitDB(); err != nil {
		t.Fatalf("DB yükelenemedi : %v", err)
		return

	}

	service := ArrsPlanningService.NewArrsPlanningService(config.Params, config.DB)

	userID := uint64(1) // Değiştir: Şifresini değiştirmek istediğin kullanıcının ID'si
	oldPassword := "112233"
	newPassword := "122333"

	changeReq := ArrsPlanningService.ChangePasswordRequest{
		UserID:      userID,
		OldPassword: oldPassword,
		NewPassword: newPassword,
	}
	if err := service.ChangeMyPassword(changeReq); err != nil {
		t.Fatalf("Şifre değiştirilirken hata oluştu: %v", err)
	} else {
		t.Log("Şifre başarıyla değiştirildi")
	}

}

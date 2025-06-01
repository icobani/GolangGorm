/*
	B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
	User        : ICI
	Name        : Ibrahim COBANI
	Date        : 18.05.2025
	Time        : 18:26
	Notes       :


*/

package Test

import (
	"log"
	"testing"

	"github.com/icobani/GolangGorm/config"
	"github.com/icobani/GolangGorm/models"
)

func TestMigrate(t *testing.T) {
	if err := config.InitConfigFile(); err != nil {
		log.Println("Error: ", err)

		return
	}

	if err := config.InitDB(); err != nil {
		log.Println("Error: ", err)
		return
	}

	if err := config.DB.AutoMigrate(&models.User{}); err != nil {
		log.Println("Error: ", err)
		return
	}

}

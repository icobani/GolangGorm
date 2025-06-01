/*
	B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
	User        : ICI
	Name        : Ibrahim COBANI
	Date        : 18.05.2025
	Time        : 17:47
	Notes       :


*/

package main

import (
	"log"

	"github.com/icobani/GolangGorm/config"
)

var (
	Version string
)

func main() {

	if err := config.InitConfigFile(); err != nil {
		log.Println("Error: ", err)
	}
	if err := config.InitDB(); err != nil {
		log.Println("Error: ", err)
	}

	InitApi()

}

/*
	B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
	User        : ICI
	Name        : Ibrahim COBANI
	Date        : 18.05.2025
	Time        : 17:52
	Notes       :


*/

package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Params  = viper.New()
	DB      *gorm.DB
	Version string
)

func InitDB() (err error) {

	cnnString := fmt.Sprintf("host=%s user=%s password='%s' dbname=%s port=%s sslmode=disable",
		Params.GetString("db.live.host"),
		Params.GetString("db.live.user"),
		Params.GetString("db.live.password"),
		Params.GetString("db.live.dbname"),
		Params.GetString("db.live.port"))

	if Params.GetString("environment") == "test" {
		cnnString = fmt.Sprintf("host=%s user=%s password='%s' dbname=%s port=%s sslmode=disable",
			Params.GetString("db.test.host"),
			Params.GetString("db.test.user"),
			Params.GetString("db.test.password"),
			Params.GetString("db.test.dbname"),
			Params.GetString("db.test.port"))
	}

	DB, err = gorm.Open(postgres.Open(cnnString), &gorm.Config{})
	if err != nil {
		log.Println("DB Error", err)
		return
	}

	log.Println("DB Connected to", Params.GetString("environment"))
	return
}

func InitConfigFile() error {
	var folderPath string
	var err error

	if Version == "" {
		// go run ile çalıştırıldığını anlıyoruz. make file'de build kodunda version set ediliyor.
		folderPath, err = os.Getwd()
		if err != nil {
			log.Println("Error: ", err)
			return err
		}

		substr := "GolangGorm"
		index := strings.Index(folderPath, substr)

		if index != -1 {
			// "api_fuel_prices" kelimesini ve öncesini almak için string'i kes
			result := folderPath[:index+len(substr)]
			folderPath = result
			fmt.Println(result)
		} else {
			fmt.Println("Substring 'GolangGorm' not found in the input.")
		}

	} else {
		folderPath, err = os.Executable()
		if err != nil {
			log.Println("Error: ", err)
			return err
		}
		folderPath = filepath.Dir(folderPath)
	}

	log.Println("Folder Path: ", folderPath)

	Params.AddConfigPath(folderPath)
	viper.SetConfigName("config.yaml")
	if err := Params.ReadInConfig(); err != nil {
		return err
	} else {
		log.Println("Config file loaded")
	}
	return nil
}

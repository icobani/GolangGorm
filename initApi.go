/*
	B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
	User        : ICI
	Name        : Ibrahim COBANI
	Date        : 1.06.2025
	Time        : 17:20
	Notes       :


*/

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitApi() {
	log.Println("Api is initializing.")
	var err error
	app := gin.Default()

	port := "8063"

	app.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	baseURL := fmt.Sprintf(":%s", port)
	err = app.Run(baseURL)
	if err != nil {
		log.Println(err)
	}
}

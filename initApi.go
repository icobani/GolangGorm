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
	"github.com/icobani/GolangGorm/apiroots"
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

	app.GET("/test", func(c *gin.Context) {

		filter := c.Query("filter")
		log.Println("Received filter query parameter:", filter)

		c.JSON(http.StatusOK, gin.H{
			"message": "test",
			"filter":  filter,
		})

	})

	app.POST("/test", func(c *gin.Context) {
		var requestData map[string]interface{}
		if err := c.BindJSON(&requestData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
			return
		}

		log.Println("Received data:", requestData)

		c.JSON(http.StatusOK, gin.H{
			"message": "test",
			"data":    requestData,
		})
	})

	app.PUT("/test", func(c *gin.Context) {
		var requestData map[string]interface{}
		if err := c.BindJSON(&requestData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
			return
		}

		log.Println("Received data:", requestData)

		c.JSON(http.StatusOK, gin.H{
			"message": "test",
			"data":    requestData,
		})
	})

	app.DELETE("/test/:id", func(c *gin.Context) {

		id := c.Param("id")
		log.Println("Received delete request for ID:", id)

		c.JSON(http.StatusOK, gin.H{
			"message": "test",
			"id":      id,
		})

	})

	api := app.Group("/api")
	apiroots.UserRoot(api)

	baseURL := fmt.Sprintf(":%s", port)
	err = app.Run(baseURL)
	if err != nil {
		log.Println(err)
	}
}

package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

var appSettings Settings

func main() {

	r := gin.Default()

	appSettings = GetSettings("settings.json")

	r.GET("/status", StatusHandler)
	r.NotFound404(NotFoundHandler)

	r.Run(":" + strconv.Itoa(appSettings.Port))
}

func NotFoundHandler(c *gin.Context) {
	c.JSON(404, gin.H{"status": "Oops, page not found"})
}

func StatusHandler(c *gin.Context) {
	c.JSON(200, gin.H{"serviceStatus": StatusMatrixGenerator(appSettings.Services)})
}

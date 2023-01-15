package server

import (
	"github.com/gin-gonic/gin"
)

// Start running the server
func Start() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/ping", ping)
	router.GET("/:code", getUrl)
	router.POST("/", genCode)

	router.Run()
}

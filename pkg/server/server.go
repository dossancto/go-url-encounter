package server

import (
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()

	router.GET("/ping", ping)
	router.POST("/", genCode)

	router.Run()
}

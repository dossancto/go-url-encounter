package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lu-css/go-url-encounter/pkg/code"
)

func ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func genCode(ctx *gin.Context) {
	code := code.GenRandomCode(6)
	ctx.JSON(http.StatusOK, gin.H{
		"genereted_code": code,
	})
}

package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lu-css/go-url-encounter/pkg/code"
	"github.com/lu-css/go-url-encounter/pkg/database"
)

// Contains the user URL.
type Url struct {
	URL string `json:"url"`
}

// Handles the classic Ping endpoint
func ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// Handles the generation POST endpoint
// Receives the client URL, save in database
// and send the code.
//
// In any case of error send a json explaining the problem.
func genCode(ctx *gin.Context) {
	var body Url
	ctx.BindJSON(&body)

	url := body.URL

	errMsg, valid := validUrl(body.URL)

	if !valid {
		ctx.JSON(http.StatusBadRequest, errMsg)
		return
	}

	code := code.GenRandomCode(6)

	err := database.CreateUrl(url, code)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Error while creting the please try again.´",
		})
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"genereted_code": code,
	})
}

func getUrl(ctx *gin.Context) {
	code := ctx.Param("code")

	url, err := database.GetUrl(code)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"code":  code,
			"error": "This code does not correspond with a URL.",
		})
		return
	}

	ctx.HTML(http.StatusOK, "redirect.tmpl", gin.H{
		"url": url,
	})

}

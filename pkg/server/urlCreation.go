package server

import (
	"github.com/gin-gonic/gin"
	"github.com/lu-css/go-url-encounter/pkg/code"
)

// Return `true` if the URL format
// is supported, for now only
// http and https are valid.
func validUrlFormat(url string) bool {
	return code.ValidProtocol(url)
}

// Valid the entire URL, using `validUrlFormat`
// function and verifies if the url is 
// empty.
func validUrl(url string) (gin.H, bool) {
	if url == "" {
		return gin.H{
			"error": "'url' field not found.",
		}, false
	}

	if !validUrlFormat(url) {
		return gin.H{
			"error": "URL protocol invalid, only http and https are supported.",
		}, false
	}

	return nil, true
}


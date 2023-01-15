package code

import (
	"regexp"
)

// Return true if the url is http or https.
// Otherwise return false.
func ValidProtocol(url string) bool {
	r, _ := regexp.Compile(`^(http|https):\/\/\S*`)
	valid := r.MatchString(url)
	return valid
}

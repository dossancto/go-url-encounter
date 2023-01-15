package code

import (
	"math/rand"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"
const numbers = "0123456789"
const special = "-._~()'!*:@,;"

const urlSafeChars = alphabet + numbers + special

// Returns a string with a random code in the specified length.
// The string if urlSafe.
func GenRandomCode(codeSize int) string {
	rand.Seed(time.Now().UnixNano())
	var resultChar string

	for i := 1; i <= codeSize; i++ {
		randChar := urlSafeChars[rand.Intn(len(urlSafeChars))]
		resultChar += string(randChar)
	}

	return resultChar
}

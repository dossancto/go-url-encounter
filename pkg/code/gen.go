package code

import (
	"math/rand"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"
const numbers = "0123456789"
const special = "-._~()'!*:@,;"

const urlSafeChars = alphabet + numbers + special

// Returns a random code with the specified length.
func GenRandomCode(codeSize int) string {
	rand.Seed(time.Now().UnixNano())
	var resultChar string

	for i := 1; i <= codeSize; i++ {
		randChar := urlSafeChars[rand.Intn(len(urlSafeChars))]
		resultChar += string(randChar)
	}

	return resultChar
}

package cobain

import (
	"math/rand"
)

// generateRandomString generates a random string of a specified length using the given characters.
const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXZ1238849103748102"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func CreateOTP() string {
	return RandStringBytes(6)
}

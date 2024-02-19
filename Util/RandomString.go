package util

import "math/rand"

const allowedChars = "_-0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GenerateRandomString(length int) string {
	state := make([]byte, length)

	for character := range state {
		state[character] = allowedChars[rand.Int63()%int64(len(allowedChars))]
	}

	return string(state)
}

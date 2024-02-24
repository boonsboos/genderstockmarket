package auth

import (
	"crypto/rand"
	"crypto/sha512"
	"math/big"
)

const secretCharacters = "qwertyuiopasdfghjklzxcvbnm0987654321QWERTYUIOPASDFGHJKLZXCVBNM-_"

func GenerateToken(username, secret string) string {

	size := big.NewInt(2 << 61)
	buf := make([]byte, 128)

	for i := 0; i < 128; i++ {
		num, _ := rand.Int(rand.Reader, size)
		buf[i] = secretCharacters[num.Int64()%int64(len(secretCharacters))]
	}

	hash := sha512.Sum512([]byte(username + secret))

	for i := range buf {
		num, _ := rand.Int(rand.Reader, big.NewInt(1000))
		if num.Int64() <= 1 {
			hashIdx, _ := rand.Int(rand.Reader, big.NewInt(63))
			buf[i] = hash[hashIdx.Int64()]
		} else {
			continue
		}
	}

	return string(buf)
}

func GenerateSecret() string {
	return GenerateRandomString(48)
}

func GenerateRandomString(length int) string {
	size := big.NewInt(2 << 61)
	buf := make([]byte, length)

	for i := 0; i < length; i++ {
		num, _ := rand.Int(rand.Reader, size)
		buf[i] = secretCharacters[num.Int64()%int64(len(secretCharacters))]
	}

	return string(buf)
}

package utils

import (
	"crypto/rand"
	"log"
)

func GenerateJWT() []byte {
	key := make([]byte, 32)

	_, err := rand.Read(key)
	if err != nil {
		log.Fatalf("Failed to generate JWT key: %v", err)
	}

	return key
}

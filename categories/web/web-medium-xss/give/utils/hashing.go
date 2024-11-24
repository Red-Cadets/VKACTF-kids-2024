package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pass string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
	return string(bytes), err

}
func CheckPasswordHash(pass, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	return err == nil
}

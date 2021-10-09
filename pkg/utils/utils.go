package utils

import "golang.org/x/crypto/bcrypt"

type Message struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

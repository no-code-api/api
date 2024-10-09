package users

import (
	"github.com/leandro-d-santos/no-code-api/internal/logger"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		log := logger.NewLogger("PasswordHashing")
		log.ErrorF("Error to generate hash: %v", err.Error())
		return "", err
	}
	return string(bytes), nil
}

func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

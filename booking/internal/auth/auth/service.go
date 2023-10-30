package auth

import (
	"booking/internal/auth/repository"
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

type Service struct {
	repo              repository.Repository
	jwtSecretKey      string
	passwordSecretKey string
	//userTransport     *transport.UserTransport
}

func (s *Service) generatePassword(password string) string {
	hash := hmac.New(sha256.New, []byte(s.passwordSecretKey))
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

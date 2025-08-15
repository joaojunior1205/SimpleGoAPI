package handlers

import (
	"bufunfa/auth"
	"errors"
	"time"
)

func LoginHandler(email string, password string) (string, error) {
	if email != "user@example.com" || password != "123456" {
		return "", errors.New("invalid credentials")
	}

	token, err := auth.GeraJWT("user-id-123", time.Hour*1)
	if err != nil {
		return "", err
	}

	return token, nil
}

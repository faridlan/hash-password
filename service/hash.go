package service

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) ([]byte, error) {

	hashesPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to generated password : %v", err)
	}

	return hashesPassword, nil

}

func ComparePasswords(hashedPassword []byte, password string) bool {

	return bcrypt.CompareHashAndPassword(hashedPassword, []byte(password)) == nil

}

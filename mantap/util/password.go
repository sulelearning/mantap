package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// membuat hashedPssword pada password user
func HashPassword(password string) (string, error) {
	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedpassword), nil
}

// compare hashedPassword pada table user dengan hashedPassword dari password user, apakah sama atau tidak
func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

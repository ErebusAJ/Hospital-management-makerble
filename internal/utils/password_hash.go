package utils

import "golang.org/x/crypto/bcrypt"

// PasswordHash
// a utility function to hash password
func PasswordHash(password string) (string, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashPass), nil
}
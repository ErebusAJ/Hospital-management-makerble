package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

// generate JWT token
func GenerateJWT(userID uuid.UUID, userRole string) (string, error) {
	expirationTime := time.Hour * 24 * 30

	claims := jwt.MapClaims{
		"userID" : userID,
		"userRole" : userRole,
		"expiry" : time.Now().Add(expirationTime).Unix(),
		"created" : time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	godotenv.Load()

	key := os.Getenv("SECRET_KEY")
	if key == "" {
		return "", errors.New("error retrieving signing key")
	}

	signedToken, err := token.SignedString([]byte(key))
	if err != nil {
		return "", errors.New("error signing key")
	}

	return signedToken, nil
}
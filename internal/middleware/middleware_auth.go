package middleware

import (
	"strings"

	"github.com/ErebusAJ/makerble-backend/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

// middleware function to handle authentication
func AuthMiddleware(tokenJWT string) gin.HandlerFunc{
	return func(c *gin.Context){
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.ErrorJSON(c, 401, "invalid header", "unable to find auth header", nil)
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return []byte(tokenJWT), nil
		})
		if err != nil{
			utils.ErrorJSON(c, 401, "invalid token", "invalid token", err)
			c.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		tempID := claims["userID"].(string)
		userRole := claims["userRole"].(string)


		userID, err := uuid.Parse(tempID)
		if err != nil {
			utils.ErrorJSON(c, 500, "internal error", "error parsing userid to uuid from body", err)
			c.Abort()
			return
		}

		c.Set("userID", userID)
		c.Set("userRole", userRole)
		c.Next()
	}
}
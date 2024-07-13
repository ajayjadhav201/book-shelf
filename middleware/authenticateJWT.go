package middleware

import (
	// "bookstore-api-go/pkg/api/admin"
	// "bookstore-api-go/pkg/api/user"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var (
	JWT_KEY = []byte(os.Getenv("SECRET_KEY"))	
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BearerSchema = "Bearer "
		header := c.GetHeader("Authorization")
		if header == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization Header"})
			c.Abort()
			return
		}

		if !strings.HasPrefix(header, BearerSchema) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization Header"})
			c.Abort()
			return
		}

		tokenStr := header[len(BearerSchema):]

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return JWT_KEY, nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// if claims.UserType != "user" {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user type"})
		// 	c.Abort()
		// 	return
		// }

		// c.Set("username", claims.Username)
		// c.Set("userType", claims.UserType)
		c.Next()
	}
}

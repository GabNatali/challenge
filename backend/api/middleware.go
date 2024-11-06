package api

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Payload struct {
	jwt.MapClaims        // ExpiryAt, IssueAt
	Session       string `json:"session"`
}

func AuthenticateSession() gin.HandlerFunc {

	return func(c *gin.Context) {

		tokenBearer := c.Request.Header.Get("Authorization")
		if tokenBearer == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No token found"})
			return
		}

		tokenParts := strings.Split(tokenBearer, "Bearer ")
		if len(tokenParts) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		token, err := jwt.ParseWithClaims(tokenParts[1], &Payload{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("invalid token")
			}
			return []byte(os.Getenv("ACCESS_TOKEN_SECRET")), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		claims, _ := token.Claims.(*Payload)

		// userData := shared.Sessions[claims.Session]
		// if userData.ExpiryTime.Before(time.Now()) {
		// 	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "expired session"})
		// 	return
		// }

		fmt.Println(claims)
		c.Next()
	}
}

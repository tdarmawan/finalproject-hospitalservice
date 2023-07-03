package middleware

// please make middleware to check token in every request
// if token is valid, continue to next handler
// if token is invalid, return error
// if token is not exist, return error
// if token is expired, return error
// Path: cmd/middleware/authentication.go
// Path Services: cmd/services/user.go
// Path handler: api/user_handler.go
// Path: cmd/middleware/authentication.go

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	pb "testing/pb/user"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CustomClaims struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func AuthMiddleware(grpcClient pb.UserDatakuClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		fmt.Println(tokenString)
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No Authorization header provided"})
			c.Abort()
			return
		}
		claims := &CustomClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization token"})
			c.Abort()
			return
		}
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, err = grpcClient.GetUserData(ctx, &pb.GetUserRequest{UserId: claims.ID})
		if err != nil {

			if status.Code(err) == codes.NotFound {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization token"})
				c.Abort()
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.Set("user_id", claims.ID)
		c.Next()
	}
}

package utils

import (
	"os"
	"time"

	"github.com/beingnoble03/octern-main/models"
	"github.com/golang-jwt/jwt"
)

func GenerateAccessTokenString(user *models.User) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(5 * time.Second).Unix(),
	})

	accessTokenString, err := accessToken.SignedString([]byte(os.Getenv("JWT_ACCESS_SECRET")))

	return accessTokenString, err
}

package utils

import (
	"os"
	"time"

	"github.com/beingnoble03/octern-main/models"
	"github.com/golang-jwt/jwt"
)

func GenerateRefreshTokenString(user *models.User) (string, error) {
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	})

	refreshTokenString, err := refreshToken.SignedString([]byte(os.Getenv("JWT_REFRESH_SECRET")))

	return refreshTokenString, err
}

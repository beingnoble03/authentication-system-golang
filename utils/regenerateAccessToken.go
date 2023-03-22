package utils

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/beingnoble03/octern-main/initializers"
	"github.com/beingnoble03/octern-main/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

func RegenerateAccessToken(c *gin.Context) {
	refreshTokenString, err := c.Cookie("RefreshToken")

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Access token and Refresh token missing.",
		})

		return
	}

	refreshToken, err := jwt.Parse(refreshTokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_REFRESH_SECRET")), nil
	})

	if err != nil {
		fmt.Println("adssad")
		c.AbortWithStatus(http.StatusBadRequest)

		return
	}

	if claims, ok := refreshToken.Claims.(jwt.MapClaims); ok && refreshToken.Valid {
		// Check if the refresh token is sent within it's expiry
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)

			return
		}

		var currentUser models.User

		result := initializers.Db.Take(&currentUser, claims["sub"])

		// User for whom the refresh token is signed doesn't exist
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusUnauthorized)

			return
		}

		// Generate new access token and set it in cookie
		accessTokenString, err := GenerateAccessTokenString(&currentUser)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Unable to generate access token.",
			})

			return
		}

		c.SetSameSite(http.SameSiteLaxMode)
		c.SetCookie("Authorization", accessTokenString, 5, "", "", false, true)

		c.Set("currentUser", currentUser)
	} else {
		// The refresh token is invalid
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Refresh token invalid.",
		})

		return
	}
}

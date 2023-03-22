package middlewares

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/beingnoble03/octern-main/initializers"
	"github.com/beingnoble03/octern-main/models"
	"github.com/beingnoble03/octern-main/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

func CheckAuth(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		utils.RegenerateAccessToken(c)

		c.Next()
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_ACCESS_SECRET")), nil
	})

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)

		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Regenerate access token using refresh token
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			utils.RegenerateAccessToken(c)

			c.Next()
		}

		var currentUser models.User

		result := initializers.Db.Take(&currentUser, claims["sub"])

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusUnauthorized)

			return
		}

		c.Set("currentUser", currentUser)

		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Authorization token invalid.",
		})

		return
	}
}

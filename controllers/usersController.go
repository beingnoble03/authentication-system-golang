package controllers

import (
	"errors"
	"net/http"

	"github.com/beingnoble03/octern-main/initializers"
	"github.com/beingnoble03/octern-main/models"
	"github.com/beingnoble03/octern-main/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateUser(c *gin.Context) {
	var data struct {
		Username       string `json:"username"`
		Password       string `json:"password"`
		OrganizationID uint   `json:"organization_id"`
	}

	if c.Bind(&data) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid parameters passed.",
		})

		return
	}

	// Checks if the currentUser is the admin of the passed organization
	currentUser, _ := c.Get("currentUser")
	var organizationUserInstance models.OrganizationUser

	result := initializers.Db.Take(&organizationUserInstance, "organization_id = ? AND user_id = ? AND is_admin = ?", data.OrganizationID, currentUser.(models.User).ID, true)

	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Current User is not an admin.",
		})

		return
	}

	// Hashes the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Something's wrong with the password. Unable to hash the password.",
		})

		return
	}

	// Creates the user instance
	user := models.User{
		Username: data.Username,
		Password: string(hashedPassword),
	}

	result = initializers.Db.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Something went wrong. Unable to create User instance.",
		})

		return
	}

	// Creates the organizationUser instance to make the user member
	organizationUser := models.OrganizationUser{
		OrganizationID: data.OrganizationID,
		UserID:         user.ID,
		IsAdmin:        false,
	}

	result = initializers.Db.Create(&organizationUser)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Something went wrong. Unable to create OrganizationUser instance.",
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":                 user.ID,
		"organizationUserId": organizationUser.ID,
	})
}

func Login(c *gin.Context) {
	var data struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if c.Bind(&data) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid parameters passed.",
		})

		return
	}

	var user models.User

	result := initializers.Db.Take(&user, "username = ?", data.Username)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid username.",
		})

		return
	}

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Unable to query database.",
		})

		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Incorrect password.",
		})

		return
	}

	// Generate new access token with 1 hour expiry time
	accessTokenString, err := utils.GenerateAccessTokenString(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Unable to generate access token.",
		})

		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", accessTokenString, 3600, "", "", false, true)

	// Generate new refresh token with 24 hours expiry time
	refreshTokenString, err := utils.GenerateRefreshTokenString(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Unable to generate refresh token.",
		})

		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("RefreshToken", refreshTokenString, 24*3600, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessTokenString,
		"refresh_token": refreshTokenString,
	})
}

func Logout(c *gin.Context) {
	c.SetCookie("Authorization", "", -1, "", "", false, true)
	c.SetCookie("RefreshToken", "", -1, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "You have been successfully logged out.",
	})
}

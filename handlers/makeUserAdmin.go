package handlers

import (
	"net/http"

	"github.com/beingnoble03/octern-main/initializers"
	"github.com/beingnoble03/octern-main/models"
	"github.com/gin-gonic/gin"
)

func MakeUserAdmin(c *gin.Context) {
	var data struct {
		UserID         uint `json:"user_id"`
		OrganizationID uint `json:"organization_id"`
	}

	if c.Bind(&data) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid parameters passed.",
		})

		return
	}

	organizationUser := models.OrganizationUser{
		UserID:         data.UserID,
		OrganizationID: data.OrganizationID,
		IsAdmin:        true,
	}

	result := initializers.Db.Create(&organizationUser)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error occured while creating OrganizatioUser instance.",
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"organizationUser": organizationUser,
	})
}

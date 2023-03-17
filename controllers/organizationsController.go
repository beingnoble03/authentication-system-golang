package controllers

import (
	"net/http"

	"github.com/beingnoble03/octern-main/initializers"
	"github.com/beingnoble03/octern-main/models"
	"github.com/gin-gonic/gin"
)

func CreateOrganization(c *gin.Context) {
	var data struct {
		Name string `json:"name"`
	}

	if c.Bind(&data) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid parameters passed.",
		})

		return
	}

	organization := models.Organization{
		Name: data.Name,
	}

	result := initializers.Db.Create(&organization)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Unable to create organization.",
		})

		return
	}

	currentUser, _ := c.Get("currentUser")

	// Make the current user the admin of the organizaiton
	organizationUser := models.OrganizationUser{
		UserID:         currentUser.(models.User).ID,
		OrganizationID: organization.ID,
		IsAdmin:        true,
	}

	result = initializers.Db.Create(&organizationUser)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error occured while declaring the user as admin.",
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": organization.ID,
	})
}

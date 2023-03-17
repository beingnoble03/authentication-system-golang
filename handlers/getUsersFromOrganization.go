package handlers

import (
	"net/http"

	"github.com/beingnoble03/octern-main/initializers"
	"github.com/beingnoble03/octern-main/models"
	"github.com/gin-gonic/gin"
)

func GetUsersFromOrganization(c *gin.Context) {
	var organizationID, _ = c.Params.Get("id")

	currentUser, _ := c.Get("currentUser")

	// Check if the organization exists
	result := initializers.Db.Where("id", organizationID).Take(&models.Organization{})

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Organization doesn't exist with the matching id.",
		})

		return
	}

	var organizationUsers []models.OrganizationUser

	result = initializers.Db.Select("user_id").Where("organization_id = ?", organizationID).Find(&organizationUsers)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error occured while fetching members.",
		})

		return
	}

	// Check if the current user is a member of the organization
	isCurrentUserMember := false

	for _, organizaitonUser := range organizationUsers {
		if organizaitonUser.UserID == currentUser.(models.User).ID {
			isCurrentUserMember = true
		}
	}

	if !isCurrentUserMember {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Only members of the organization can view other members.",
		})

		return
	}

	var userIDs []uint

	for _, organizationUser := range organizationUsers {
		userIDs = append(userIDs, organizationUser.UserID)
	}

	var users []models.User

	result = initializers.Db.Select("id", "username").Find(&users, userIDs)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error occured while fetching users",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"members":         users,
		"organization_id": organizationID,
	})
}

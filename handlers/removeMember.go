package handlers

import (
	"net/http"

	"github.com/beingnoble03/octern-main/initializers"
	"github.com/beingnoble03/octern-main/models"
	"github.com/gin-gonic/gin"
)

func RemoveMember(c *gin.Context) {
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

	// Finds the organizatioUser instance for the desired user and organization.
	var organizationUser models.OrganizationUser

	result = initializers.Db.Take(&organizationUser, "organization_id = ? AND user_id = ?", data.OrganizationID, data.UserID)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "The user is not a member of the organization.",
		})

		return
	}

	result = initializers.Db.Delete(&organizationUser)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Unable to remove member from the organization.",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

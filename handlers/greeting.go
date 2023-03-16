package handlers

import (
	"net/http"

	"github.com/beingnoble03/octern-main/models"
	"github.com/gin-gonic/gin"
)

func Greeting(c *gin.Context) {
	currentUser, _ := c.Get("currentUser")
	currentUserUsername := currentUser.(models.User).Username

	message := "Hi! This is a Go Project."

	c.JSON(http.StatusOK, gin.H{
		"message":  message,
		"username": currentUserUsername,
	})
}

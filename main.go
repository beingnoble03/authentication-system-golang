package main

import (
	"fmt"

	"github.com/beingnoble03/octern-main/controllers"
	"github.com/beingnoble03/octern-main/handlers"
	"github.com/beingnoble03/octern-main/initializers"
	"github.com/beingnoble03/octern-main/middlewares"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDb()
	fmt.Println("Initiated")
}

func main() {
	r := gin.Default()

	// These endpoints are not in the scope of task. Made for easier modification of data.
	r.GET("/validate", middlewares.CheckAuth, handlers.Greeting)
	r.POST("/createOrganization", controllers.CreateOrganization)
	r.POST("/makeUserAdmin", handlers.MakeUserAdmin)

	// These endpoints are in the scope of task.
	r.GET("/getUsersFromOrganization/:id", middlewares.CheckAuth, handlers.GetUsersFromOrganization)
	r.POST("/createUser", middlewares.CheckAuth, controllers.CreateUser)
	r.POST("/login", controllers.Login)
	r.POST("/removeMember", middlewares.CheckAuth, handlers.RemoveMember)
	r.GET("/logout", middlewares.CheckAuth, controllers.Logout)

	r.Run()
}

package main

import (
	"Assignment2/controller"
	"Assignment2/middleware"
	"Assignment2/model"

	"github.com/gin-gonic/gin"
)

// @title           Assignment Golang 2
// @version         1.0
// @description     This is a REST API server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   Brian Habib
// @contact.email  brianhabib252@gmail.com
// @host      localhost:8080
// @BasePath  /Assignment2

func main() {
	model.StartdB()

	r := gin.Default()

	// Database Routes
	r.GET("/orders/:id", controller.GetOrder)
	r.GET("/orders", controller.GetAllOrders)
	r.POST("/orders", middleware.Authentication, controller.CreateOrder)
	r.PUT("/orders/:id", middleware.Authentication, controller.UpdateOrder)
	r.DELETE("/orders/:id", middleware.Authentication, controller.DeleteOrder)

	// Authentication Routes
	r.POST("/signup", controller.SignUp)
	r.POST("/signin", controller.SignIn)
	r.POST("/signout", controller.SignOut)

	// Run the server
	r.Run(":8080")
}

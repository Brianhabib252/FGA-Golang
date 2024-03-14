package main

import (
	//"Assignment2/controller"
	"Assignment2/controller"
	"Assignment2/model"

	//"github.com/swaggo/http-swagger"
	//"fmt"
	"github.com/gin-gonic/gin"
	//"gorm.io/driver/mysql"
	//"gorm.io/gorm"
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

	// Routes
	r.GET("/orders/:id", controller.GetOrder)
	r.GET("/orders", controller.GetAllOrders)
	r.POST("/orders", controller.CreateOrder)
	r.PUT("/orders/:id", controller.UpdateOrder)
	r.DELETE("/orders/:id", controller.DeleteOrder)

	// Run the server
	r.Run(":8080")
}

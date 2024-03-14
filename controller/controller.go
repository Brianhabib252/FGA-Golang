package controller

import (
	"Assignment2/model"

	"github.com/gin-gonic/gin"
	//"gorm.io/gorm"
)

// ShowAccount godoc
// @Summary      get an order
// @Description  get order by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200  {object}  model.Account
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /accounts/{id} [get]
func GetOrder(c *gin.Context) {
	db := model.GetDB()
	id := c.Param("id")
	var order model.Order
	if err := db.Preload("Items").First(&order, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Order not found"})
		return
	}
	c.JSON(200, order)
}

// ShowAccount godoc
// @Summary      get all data
// @Description  get all data
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200  {object}  model.Account
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /accounts/{id} [get]
func GetAllOrders(c *gin.Context) {
	db := model.GetDB()
	var orders []model.Order
	if err := db.Preload("Items").Find(&orders).Error; err != nil {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(200, orders)
}

// ShowAccount godoc
// @Summary      crate a new data
// @Description  crate a new data
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200  {object}  model.Account
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /accounts/{id} [get]
func CreateOrder(c *gin.Context) {
	db := model.GetDB()
	var request model.Order
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	db.Create(&request)
	c.JSON(201, gin.H{"message": "Order created successfully", "order": request})
}

// ShowAccount godoc
// @Summary      update order data
// @Description  update order data by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200  {object}  model.Account
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /accounts/{id} [get]
func UpdateOrder(c *gin.Context) {
	db := model.GetDB()
	id := c.Param("id")
	var order model.Order
	if err := db.First(&order, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Order not found"})
		return
	}
	var request model.Order
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	db.Model(&order).Updates(request)
	c.JSON(200, gin.H{"message": "Order updated successfully", "order": order})
}

// ShowAccount godoc
// @Summary      delete order data
// @Description  delete order data by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200  {object}  model.Account
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /accounts/{id} [get]
func DeleteOrder(c *gin.Context) {
	db := model.GetDB()
	id := c.Param("id")
	var order model.Order
	if err := db.First(&order, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Order not found"})
		return
	}
	db.Delete(&order)
	c.JSON(200, gin.H{"message": "Order deleted successfully"})
}

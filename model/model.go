package model

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// user model
type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
}

// Item model
type Item struct {
	ID          uint   `gorm:"primaryKey",json:"lineItemId"`
	ItemCode    string `json:"itemCode"`
	Description string
	Quantity    int
	OrderID     int `json:"OrderId"`
}

// Order model
type Order struct {
	ID            uint      `gorm:"primaryKey"`
	Customer_name string    `json:"CustomerName"`
	OrderedAt     time.Time `json:"OrderedAt"`
	Items         []Item    `gorm:"foreignKey:OrderID"`
}

var db *gorm.DB

func StartdB() {
	dsn := "root:Purnama02@tcp(127.0.0.1:3306)/Assignment2?parseTime=true"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return
	}

	// Create tables
	if err := db.AutoMigrate(&Order{}, &Item{}); err != nil {
		fmt.Println("Failed to migrate tables:", err)
		return
	}
	db.AutoMigrate(&User{})

}

func GetDB() *gorm.DB {
	return db
}

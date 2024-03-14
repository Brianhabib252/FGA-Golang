package model

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Item model
type Item struct {
	ID          uint   `gorm:"primaryKey"`
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
	Items         []Item    `gorm:"foreignKey:OrderID",json:"items"`
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

}

func GetDB() *gorm.DB {
	return db
}

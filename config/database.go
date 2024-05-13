package config

import (
	"booking-api/models"

	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func InitDB() {
	dns := os.Getenv("MYSQL_DNS")
	DB, err = gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(&models.Customer{}, &models.Invoice{}, &models.Vehicle{}, &models.Employee{}, &models.Booking{})
}

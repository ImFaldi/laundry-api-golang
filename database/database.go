package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Conn() {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/laundry?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	DB = database
}

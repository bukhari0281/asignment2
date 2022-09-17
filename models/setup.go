package models

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
)

func ConnectDatabase() *gorm.DB {
	dsn := "user:pass@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to database : error=%v", err)
		return nil
	}

	return db

}

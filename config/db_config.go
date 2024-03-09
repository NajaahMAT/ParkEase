package config

import (
	"ParkEase/helper"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	user     = "root"       // change as per your MySQL user
	password = "Najaah@123" // change as per your MySQL password
	dbName   = "park_ease"
	host     = "localhost"
	port     = 3306
)

func DatabaseConnection() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	helper.ErrorPanic(err)

	return db
}

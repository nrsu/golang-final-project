package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() {
	dsn := "root:boxer312004@tcp(host.docker.internal:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = database
}

func GetDB() *gorm.DB {
	return db
}

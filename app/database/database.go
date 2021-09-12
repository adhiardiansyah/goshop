package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func GetDB() (*gorm.DB, error) {
	DB_USER := "root"
	DB_PASS := ""
	DB_HOST := "localhost"
	DB_NAME := "goshop"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", DB_USER, DB_PASS, DB_HOST, DB_NAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

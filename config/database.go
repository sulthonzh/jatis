package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewDB() (*gorm.DB, error) {
	host := "localhost"
	port := "3306"
	dbname := "jatis"
	username := "root"
	password := "root"

	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8&parseTime=true&loc=Local"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
}

package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var con *gorm.DB

func connect() *gorm.DB {
	dsn := "your_username:your_password@tcp(localhost:3306)/your_database?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func GetDBConnection() *gorm.DB {
	if con == nil {
		con = connect()
	}
	return con
}

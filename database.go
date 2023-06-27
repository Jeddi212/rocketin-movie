package main

import (
	"rocketin-movie/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var con *gorm.DB

func connect() *gorm.DB {
	dsn := "host=localhost user=jeddi dbname=moviedb port=5432 TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
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

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Movie{})
}

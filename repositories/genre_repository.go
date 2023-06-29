package repositories

import (
	"rocketin-movie/models"

	"gorm.io/gorm"
)

func FindGenreByName(db *gorm.DB, name string) (models.Genre, error) {
	var genre models.Genre
	err := db.Where("name = ?", name).First(&genre).Error
	return genre, err
}

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

func CreateGenre(db *gorm.DB, genre models.Genre) (models.Genre, error) {
	err := db.Create(&genre).Error
	return genre, err
}

func IncrementGenreViews(db *gorm.DB, genre models.Genre) error {
	genre.ViewCount += 1
	return db.Save(&genre).Error
}

func DecrementGenreViews(db *gorm.DB, genre models.Genre) error {
	genre.ViewCount -= 1
	return db.Save(&genre).Error
}

func FetchMostViewedGenre(db *gorm.DB) ([]models.Genre, error) {
	var genre []models.Genre
	result := db.Raw("SELECT * FROM genres WHERE view_count = (" +
		"SELECT MAX(view_count)" +
		"FROM genres" +
		")").Scan(&genre)
	return genre, result.Error
}

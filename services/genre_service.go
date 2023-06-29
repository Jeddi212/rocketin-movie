package services

import (
	"rocketin-movie/models"
	"rocketin-movie/repositories"

	"gorm.io/gorm"
)

func GetGenres(db *gorm.DB, names []models.GenreDTO) []models.Genre {
	var genres []models.Genre

	for _, genre := range names {
		g, _ := repositories.FindGenreByName(db, genre.Name)

		genres = append(genres, g)
	}

	return genres
}

package services

import (
	"rocketin-movie/models"
	"rocketin-movie/repositories"
	"strings"

	"gorm.io/gorm"
)

func GetGenres(db *gorm.DB, n string) []models.Genre {
	names := strings.Split(n, "")
	var genres []models.Genre

	for _, name := range names {
		g, _ := repositories.FindGenreByName(db, name)

		genres = append(genres, g)
	}

	return genres
}

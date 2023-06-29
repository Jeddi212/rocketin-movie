package services

import (
	"rocketin-movie/models"
	"rocketin-movie/repositories"

	"gorm.io/gorm"
)

func GetGenres(db *gorm.DB, names []models.GenreDTO) ([]models.Genre, error) {
	var genres []models.Genre

	for _, genre := range names {
		g, err := repositories.FindGenreByName(db, genre.Name)

		if err != nil {
			return nil, err
		}

		genres = append(genres, g)
	}

	return genres, nil
}

func CreateNewGenre(db *gorm.DB, dto models.GenreDTO) (models.Genre, error) {
	return repositories.CreateGenre(db, GenreCreateMapper(dto))
}

func GenreCreateMapper(dto models.GenreDTO) models.Genre {
	return models.Genre{
		Name: dto.Name,
	}
}

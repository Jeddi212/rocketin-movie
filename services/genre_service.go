package services

import (
	"rocketin-movie/models"
	"rocketin-movie/models/dto"
	"rocketin-movie/repositories"

	"gorm.io/gorm"
)

func GetGenres(db *gorm.DB, names []dto.GenreDTO) ([]models.Genre, error) {
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

func CreateNewGenre(db *gorm.DB, dto dto.GenreDTO) (models.Genre, error) {
	return repositories.CreateGenre(db, GenreCreateMapper(dto))
}

func WatchGenre(db *gorm.DB, genres []models.Genre) error {
	for _, genre := range genres {
		genre, _ := repositories.FindGenreByName(db, genre.Name)
		err := repositories.IncrementGenreViews(db, genre)

		if err != nil {
			return err
		}
	}

	return nil
}

func GenreCreateMapper(dto dto.GenreDTO) models.Genre {
	return models.Genre{
		Name: dto.Name,
	}
}

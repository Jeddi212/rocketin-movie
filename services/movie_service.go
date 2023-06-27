package services

import (
	"gorm.io/gorm"

	"rocketin-movie/models"
	"rocketin-movie/repositories"
)

func SearchMovies(db *gorm.DB, movieDTO models.MovieDTO) []models.Movie {
	if movieDTO.Title == "" && movieDTO.Description == "" && movieDTO.Artists == "" && movieDTO.Genres == "" {
		return repositories.FetchAllMovies(db)
	}
	return repositories.FindMovies(db, movieDTO)
}

func UpdateMovie(db *gorm.DB, movieID string, movieDTO models.Movie) error {
	movie, err := repositories.FindMovieByID(db, movieID)
	if err != nil {
		return err
	}
	return repositories.UpdateMovie(db, movie, movieDTO)
}

func CreateNewMovie(db *gorm.DB, movie models.Movie) (models.Movie, error) {
	return repositories.CreateMovie(db, movie)
}

func WatchMovie(db *gorm.DB, movieID string) error {
	movie, err := repositories.FindMovieByID(db, movieID)
	if err != nil {
		return err
	}
	return repositories.IncrementWatchNumber(db, movie, movieID)
}

package services

import (
	"gorm.io/gorm"

	"rocketin-movie/models"
	"rocketin-movie/repositories"
)

func SearchMovies(db *gorm.DB, dto models.MovieSearchDTO) []models.Movie {
	if dto.Title == "" && dto.Description == "" && dto.Artists == "" && len(dto.Genres) == 0 {
		return repositories.FetchAllMovies(db)
	}
	return repositories.FindMovies(db, dto)
}

func CreateNewMovie(db *gorm.DB, dto models.MovieCreateDTO) (models.Movie, error) {
	genres, err := GetGenres(db, dto.Genres)
	if err != nil {
		return models.Movie{}, err
	}

	movie := MovieCreateMapper(dto, genres)
	return repositories.CreateMovie(db, movie)
}

func UpdateMovie(db *gorm.DB, movieID string, dto models.MovieCreateDTO) (models.Movie, error) {
	movie, err := repositories.FindMovieByID(db, movieID)
	if err != nil {
		return models.Movie{}, err
	}

	genres, err := GetGenres(db, dto.Genres)
	if err != nil {
		return models.Movie{}, err
	}

	movie = MovieUpdateMapper(dto, movie, genres)
	return repositories.UpdateMovie(db, movie)
}

func WatchMovie(db *gorm.DB, movieID string) error {
	movie, err := repositories.FindMovieByID(db, movieID)
	if err != nil {
		return err
	}
	return repositories.IncrementWatchNumber(db, movie, movieID)
}

func MovieCreateMapper(dto models.MovieCreateDTO, genres []models.Genre) models.Movie {
	var movie models.Movie

	if dto.Title != "" {
		movie.Title = dto.Title
	}

	if dto.Description != "" {
		movie.Description = dto.Description
	}

	if dto.Duration != 0 {
		movie.Duration = dto.Duration
	}

	if dto.Artists != "" {
		movie.Artists = dto.Artists
	}

	if len(genres) != 0 {
		movie.Genres = genres
	}

	if dto.WatchURL != "" {
		movie.WatchURL = dto.WatchURL
	}

	return movie
}

func MovieUpdateMapper(dto models.MovieCreateDTO, movie models.Movie, genres []models.Genre) models.Movie {
	if dto.Title != "" {
		movie.Title = dto.Title
	}

	if dto.Description != "" {
		movie.Description = dto.Description
	}

	if dto.Duration != 0 {
		movie.Duration = dto.Duration
	}

	if dto.Artists != "" {
		movie.Artists = dto.Artists
	}

	if len(genres) != 0 {
		movie.Genres = genres
	}

	if dto.WatchURL != "" {
		movie.WatchURL = dto.WatchURL
	}

	return movie
}

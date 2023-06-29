package services

import (
	"fmt"

	"gorm.io/gorm"

	"rocketin-movie/models"
	"rocketin-movie/repositories"
)

func SearchMovies(db *gorm.DB, dto models.MovieSearchDTO) []models.Movie {
	if dto.Title == "" && dto.Description == "" && dto.Artists == "" && dto.Genres == "" {
		return repositories.FetchAllMovies(db)
	}
	return repositories.FindMovies(db, dto)
}

func CreateNewMovie(db *gorm.DB, dto models.MovieCreateDTO) (models.Movie, error) {
	movie := MovieCreateMapper(dto, GetGenres(db, dto.Genres))
	return repositories.CreateMovie(db, movie)
}

func UpdateMovie(db *gorm.DB, movieID string, dto models.MovieCreateDTO) error {
	movie, err := repositories.FindMovieByID(db, movieID)
	if err != nil {
		return err
	}

	movie = MovieUpdateMapper(dto, movie, GetGenres(db, dto.Genres))
	return repositories.UpdateMovie(db, movie)
}

func WatchMovie(db *gorm.DB, movieID string) error {
	movie, err := repositories.FindMovieByID(db, movieID)
	fmt.Println("Sampe sini")
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

	if dto.Artists != "" {
		movie.Artists = dto.Artists
	}

	if len(genres) == 0 {
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

	if dto.Artists != "" {
		movie.Artists = dto.Artists
	}

	if len(genres) == 0 {
		movie.Genres = genres
	}

	if dto.WatchURL != "" {
		movie.WatchURL = dto.WatchURL
	}

	return movie
}

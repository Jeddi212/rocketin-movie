package services

import (
	"gorm.io/gorm"

	"rocketin-movie/models"
	"rocketin-movie/models/dto"
	"rocketin-movie/models/extra"
	"rocketin-movie/repositories"
)

func SearchAllMovies(db *gorm.DB, pagination extra.Pagination) []models.Movie {
	offset, limit := Paginate(pagination)
	return repositories.FetchAllMovies(db, offset, limit)
}

func SearchMovies(db *gorm.DB, dto dto.MovieSearchDTO) []models.Movie {
	return repositories.FindMovies(db, dto)
}

func CreateNewMovie(db *gorm.DB, dto dto.MovieCreateDTO) (models.Movie, error) {
	genres, err := GetGenres(db, dto.Genres)
	if err != nil {
		return models.Movie{}, err
	}

	movie := MovieCreateMapper(dto, genres)
	return repositories.CreateMovie(db, movie)
}

func UpdateMovie(db *gorm.DB, movieID string, dto dto.MovieCreateDTO) (models.Movie, error) {
	movie, err := repositories.FindMovieByID(db, movieID)

	if err != nil {
		return models.Movie{}, err
	}

	err = repositories.ClearMovieGenre(db, movie.ID)
	if err != nil {
		return models.Movie{}, err
	}

	err = DecrementGenreViews(db, movie.Genres)
	if err != nil {
		return models.Movie{}, err
	}

	genres, err := GetGenres(db, dto.Genres)
	if err != nil {
		return models.Movie{}, err
	}

	err = IncrementGenreViews(db, genres)
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

	err = WatchGenre(db, movie.Genres)
	if err != nil {
		return err
	}
	return repositories.IncrementWatchNumber(db, movie)
}

func MovieCreateMapper(dto dto.MovieCreateDTO, genres []models.Genre) models.Movie {
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

func MovieUpdateMapper(dto dto.MovieCreateDTO, movie models.Movie, genres []models.Genre) models.Movie {
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

func Paginate(pagination extra.Pagination) (int, int) {
	offset, limit := (pagination.Page-1)*pagination.Limit, pagination.Limit

	if offset < 1 {
		offset = -1
	}

	if limit < 1 {
		limit = 10
	}

	return offset, limit
}

package repositories

import (
	"rocketin-movie/models"
	"rocketin-movie/models/dto"

	"strings"

	"gorm.io/gorm"
)

func FetchAllMovies(db *gorm.DB, offset int, limit int) []models.Movie {
	var movies []models.Movie
	db.Preload("Genres").Offset(offset).Limit(limit).Find(&movies)
	return movies
}

func FindMovies(db *gorm.DB, term dto.MovieSearchDTO) []models.Movie {
	var movies []models.Movie
	db.Preload("Genres").Where("LOWER(title) LIKE ? OR LOWER(description) LIKE ? OR LOWER(artists) LIKE ?",
		"%"+strings.ToLower(term.Title)+"%",
		"%"+strings.ToLower(term.Description)+"%",
		"%"+strings.ToLower(term.Artists)+"%").Find(&movies)

	return movies
}

func FindMovieByID(db *gorm.DB, movieID string) (models.Movie, error) {
	var movie models.Movie
	result := db.Preload("Genres").First(&movie, movieID)
	return movie, result.Error
}

func CreateMovie(db *gorm.DB, movie models.Movie) (models.Movie, error) {
	db.Model(&movie).Association("Genres")
	result := db.Create(&movie)
	return movie, result.Error
}

func UpdateMovie(db *gorm.DB, movie models.Movie) (models.Movie, error) {
	result := db.Save(&movie)
	return movie, result.Error
}

func ClearMovieGenre(db *gorm.DB, movieID uint) error {
	return db.Unscoped().Exec("DELETE FROM movie_genres WHERE movie_id = ?", movieID).Error
}

func IncrementWatchNumber(db *gorm.DB, movie models.Movie) error {
	movie.Watch += 1
	return db.Save(&movie).Error
}

func FetchMostViewedMovie(db *gorm.DB) ([]models.Movie, error) {
	var movie []models.Movie
	result := db.Raw("SELECT * FROM movies WHERE watch = (" +
		"SELECT MAX(watch)" +
		"FROM movies" +
		")").Scan(&movie)
	return movie, result.Error
}

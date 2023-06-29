package repositories

import (
	"rocketin-movie/models"

	"gorm.io/gorm"
)

func FetchAllMovies(db *gorm.DB) []models.Movie {
	var movies []models.Movie

	db.First(&movies)

	return movies
}

func FindMovies(db *gorm.DB, term models.MovieSearchDTO) []models.Movie {
	var movies []models.Movie

	db.Where("title LIKE ? OR description LIKE ? OR artists LIKE ? OR genres LIKE ?",
		"%"+term.Title+"%", "%"+term.Description+"%", "%"+term.Artists+"%").Find(&movies)
	// "%"+term.Title+"%", "%"+term.Description+"%", "%"+term.Artists+"%", "%"+term.Genres+"%").Find(&movies)

	return movies
}

func FindMovieByID(db *gorm.DB, movieID string) (models.Movie, error) {
	var movie models.Movie
	result := db.First(&movie, movieID)
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

func IncrementWatchNumber(db *gorm.DB, movie models.Movie, movieID string) error {
	movie.Watch += 1
	result := db.Save(&movie)
	return result.Error
}

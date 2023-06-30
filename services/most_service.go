package services

import (
	"rocketin-movie/models"
	"rocketin-movie/models/extra"
	"rocketin-movie/repositories"

	"gorm.io/gorm"
)

func SearchMostViewed(db *gorm.DB) (extra.MostViewed, error) {
	movie, err := repositories.FetchMostViewedMovie(db)
	if err != nil {
		return extra.MostViewed{}, err
	}

	genre, err := repositories.FetchMostViewedGenre(db)
	if err != nil {
		return extra.MostViewed{}, err
	}

	return MostViewedMapper(movie, genre), nil
}

func MostViewedMapper(movie []models.Movie, genre []models.Genre) extra.MostViewed {
	return extra.MostViewed{
		Movie: movie,
		Genre: genre,
	}
}

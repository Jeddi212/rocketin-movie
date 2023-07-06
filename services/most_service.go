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

func ListMostVotedMovieAndViewedGenre(db *gorm.DB) (extra.MostVotedMovieAndViewedGenre, error) {
	mostVotedMovie, err := repositories.FetchMostVotesMovies(db)
	if err != nil {
		return extra.MostVotedMovieAndViewedGenre{}, err
	}

	movies, err := repositories.FindMovieByIds(db, GetMovieIDsFromMostVoted(mostVotedMovie))
	if err != nil {
		return extra.MostVotedMovieAndViewedGenre{}, err
	}

	genres, err := repositories.FetchMostViewedGenre(db)
	if err != nil {
		return extra.MostVotedMovieAndViewedGenre{}, err
	}

	return extra.MostVotedMovieAndViewedGenre{
		Votes:  mostVotedMovie[0].VoteCount,
		Movies: movies,
		Genres: genres,
	}, err
}

func MostViewedMapper(movie []models.Movie, genre []models.Genre) extra.MostViewed {
	return extra.MostViewed{
		Movie: movie,
		Genre: genre,
	}
}

func GetMovieIDsFromMostVoted(movieVotes []extra.MostVotedMovie) []int {
	movieIDs := make([]int, len(movieVotes))
	for i, mv := range movieVotes {
		movieIDs[i] = mv.MovieID
	}
	return movieIDs
}

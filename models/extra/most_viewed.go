package extra

import "rocketin-movie/models"

type MostViewed struct {
	Movie []models.Movie `json:"movie"`
	Genre []models.Genre `json:"genre"`
}

type MostVotedMovie struct {
	MovieID   int `json:"movie_id"`
	VoteCount int `json:"vote_count"`
}

type MostVotedMovieAndViewedGenre struct {
	Votes  int            `json:"votes"`
	Movies []models.Movie `json:"movies"`
	Genres []models.Genre `json:"genres"`
}

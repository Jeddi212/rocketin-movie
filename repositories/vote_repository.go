package repositories

import (
	"rocketin-movie/models/extra"

	"gorm.io/gorm"
)

func AppendVote(db *gorm.DB, movieID string, username string) error {
	return db.Exec("INSERT INTO votes (user_username, movie_id) VALUES (?, ?)", username, movieID).Error
}

func DownVote(db *gorm.DB, movieID string, username string) error {
	return db.Unscoped().Exec("DELETE FROM votes WHERE user_username = ? AND movie_id = ?", username, movieID).Error
}

func FetchUserVotes(db *gorm.DB, username string) ([]int, error) {
	var movieIds []int
	result := db.Table("votes").Select("movie_id").Where("user_username = ?", username).Find(&movieIds)
	return movieIds, result.Error
}

func FetchMostVotesMovies(db *gorm.DB) ([]extra.MostVotedMovie, error) {
	var mostVoted []extra.MostVotedMovie
	result := db.Raw("SELECT movie_id, vote_count FROM (SELECT movie_id, COUNT(*) AS vote_count " +
		"FROM votes GROUP BY movie_id) AS counts WHERE vote_count = (SELECT MAX(vote_count) " +
		"FROM (SELECT COUNT(*) AS vote_count FROM votes GROUP BY movie_id) AS subcounts)").
		Scan(&mostVoted)
	return mostVoted, result.Error
}

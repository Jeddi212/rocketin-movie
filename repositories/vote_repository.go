package repositories

import "gorm.io/gorm"

func AppendVote(db *gorm.DB, movieID string, username string) error {
	return db.Exec("INSERT INTO votes (user_username, movie_id) VALUES (?, ?)", username, movieID).Error
}

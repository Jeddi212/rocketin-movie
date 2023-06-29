package models

import "gorm.io/gorm"

type Votes struct {
	gorm.Model
	MovieID uint   `json:"movieID"`
	UserID  uint   `json:"userID"`
	Action  string `json:"action"` // UPVOTE(1) DOWNVOTE(-1)
}

package services

import (
	"errors"
	"rocketin-movie/models"
	"rocketin-movie/models/dto"
	"rocketin-movie/repositories"

	"gorm.io/gorm"
)

func UpvoteMovie(db *gorm.DB, movieID string, dto dto.VoteDTO) (models.Movie, error) {
	movie, err := repositories.FindMovieByID(db, movieID)
	if err != nil {
		return models.Movie{}, err
	}

	user, err := repositories.FindUserById(db, dto.Username)
	if err != nil {
		return models.Movie{}, err
	}

	if user.Level < 1 {
		return models.Movie{}, errors.New("your membership level can't vote")
	}

	err = repositories.AppendVote(db, movieID, dto.Username)
	if err != nil {
		return models.Movie{}, err
	}

	return movie, nil
}

func DownVoteMovie(db *gorm.DB, movieID string, dto dto.VoteDTO) (models.Movie, error) {
	movie, err := repositories.FindMovieByID(db, movieID)
	if err != nil {
		return models.Movie{}, err
	}

	user, err := repositories.FindUserById(db, dto.Username)
	if err != nil {
		return models.Movie{}, err
	}

	if user.Level < 1 {
		return models.Movie{}, errors.New("your membership level can't downvote")
	}

	err = repositories.DownVote(db, movieID, dto.Username)
	if err != nil {
		return models.Movie{}, err
	}

	return movie, nil
}

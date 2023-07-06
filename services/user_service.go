package services

import (
	"rocketin-movie/models"
	"rocketin-movie/models/dto"
	"rocketin-movie/repositories"

	"gorm.io/gorm"
)

func RegisterUser(db *gorm.DB, dto dto.UserDTO) error {
	return repositories.AddNewMember(db, UserMapper(dto))
}

func UserMapper(dto dto.UserDTO) models.User {
	return models.User{
		Username: dto.Username,
		Password: dto.Password,
		Level:    1,
		Votes:    []models.Movie{},
	}
}

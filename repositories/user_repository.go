package repositories

import (
	"rocketin-movie/models"

	"gorm.io/gorm"
)

func AddNewMember(db *gorm.DB, user models.User) error {
	return db.Create(&user).Error
}

func FindUserById(db *gorm.DB, username string) (models.User, error) {
	var user models.User
	result := db.First(&user, "username = ?", username)
	return user, result.Error
}

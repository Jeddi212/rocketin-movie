package repositories

import (
	"rocketin-movie/models"

	"gorm.io/gorm"
)

func AddNewMember(db *gorm.DB, user models.User) error {
	return db.Create(&user).Error
}

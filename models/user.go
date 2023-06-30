package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string  `json:"name"`
	Level int     `json:"level"`
	Votes []Movie `json:"votes" gorm:"many2many:votes;"`
}

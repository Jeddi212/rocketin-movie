package models

type Genre struct {
	Name      string `json:"name" gorm:"primaryKey"`
	ViewCount int    `json:"viewCount"`
}

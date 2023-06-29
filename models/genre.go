package models

type Genre struct {
	Name string `json:"name" gorm:"primaryKey"`
	View string `json:"view"`
}

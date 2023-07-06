package models

type User struct {
	Username string  `json:"username" gorm:"primaryKey"`
	Password string  `json:"password"`
	Level    int     `json:"level"`
	Votes    []Movie `json:"votes" gorm:"many2many:votes;"`
}

package models

type MovieDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Artists     string `json:"artists"`
	Genres      string `json:"genres"`
}

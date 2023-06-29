package dto

type MovieSearchDTO struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Artists     string     `json:"artists"`
	Genres      []GenreDTO `json:"genres"`
}

type MovieCreateDTO struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Duration    int        `json:"duration"`
	Artists     string     `json:"artists"`
	Genres      []GenreDTO `json:"genres"`
	Watch       int        `json:"watch"`
	WatchURL    string     `json:"watch_url"`
}

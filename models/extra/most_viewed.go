package extra

import "rocketin-movie/models"

type MostViewed struct {
	Movie []models.Movie `json:"movie"`
	Genre []models.Genre `json:"genre"`
}

package controllers

import (
	"encoding/json"
	"net/http"

	"rocketin-movie/models"
	"rocketin-movie/services"

	"gorm.io/gorm"
)

type GenreController struct {
	DB *gorm.DB
}

func NewGenreController(db *gorm.DB) *GenreController {
	return &GenreController{DB: db}
}

func (mc *GenreController) CreateGenre(w http.ResponseWriter, r *http.Request) {
	var dto models.GenreDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	genre, err := services.CreateNewGenre(mc.DB, dto)
	var response models.Response
	if err != nil {
		response = models.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to add a new genre",
			Data:       err,
		}
	} else {
		response = models.Response{
			StatusCode: http.StatusOK,
			Message:    "Success adding a new genre",
			Data:       genre,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

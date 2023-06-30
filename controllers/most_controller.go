package controllers

import (
	"encoding/json"
	"net/http"

	"rocketin-movie/models/extra"
	"rocketin-movie/services"

	"gorm.io/gorm"
)

type MostController struct {
	DB *gorm.DB
}

func NewMostController(db *gorm.DB) *MostController {
	return &MostController{DB: db}
}

func (mc *MostController) GetMostViewed(w http.ResponseWriter, r *http.Request) {
	mostViewed, err := services.SearchMostViewed(mc.DB)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := extra.Response{
		StatusCode: http.StatusOK,
		Message:    "Success get the most viewed movie and genre",
		Data:       mostViewed,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

package controllers

import (
	"encoding/json"
	"net/http"
	"rocketin-movie/models/dto"
	"rocketin-movie/models/extra"
	"rocketin-movie/services"

	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{DB: db}
}

func (uc *UserController) Register(w http.ResponseWriter, r *http.Request) {
	var dto dto.UserDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = services.RegisterUser(uc.DB, dto)
	var response extra.Response
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		response = extra.Response{
			StatusCode: http.StatusOK,
			Message:    "Successfully register to Rocketin Movie",
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

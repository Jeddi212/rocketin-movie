package controllers

import (
	"encoding/json"
	"net/http"

	"rocketin-movie/models"
	"rocketin-movie/services"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type MovieController struct {
	DB *gorm.DB
}

func NewMovieController(db *gorm.DB) *MovieController {
	return &MovieController{DB: db}
}

func (mc *MovieController) GetMovie(w http.ResponseWriter, r *http.Request) {
	var dto models.MovieSearchDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var movie []models.Movie = services.SearchMovies(mc.DB, dto)
	var response models.Response = models.Response{
		StatusCode: http.StatusOK,
		Message:    "Success adding a new movie",
		Data:       movie,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (mc *MovieController) CreateMovie(w http.ResponseWriter, r *http.Request) {
	var dto models.MovieCreateDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	movie, errs := services.CreateNewMovie(mc.DB, dto)
	var response models.Response
	if errs != nil {
		response = models.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to update Movie with ID",
			Data:       errs,
		}
	} else {
		response = models.Response{
			StatusCode: http.StatusOK,
			Message:    "Success adding a new movie",
			Data:       movie,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (mc *MovieController) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	var dto models.MovieCreateDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	movieID := mux.Vars(r)["movie_id"]

	err = services.UpdateMovie(mc.DB, movieID, dto)
	var response models.Response
	if err != nil {
		response = models.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to update Movie with ID " + movieID,
			Data:       err,
		}
	} else {
		response = models.Response{
			StatusCode: http.StatusOK,
			Message:    "Success update movie with ID " + movieID,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (mc *MovieController) WatchMovie(w http.ResponseWriter, r *http.Request) {
	movieID := mux.Vars(r)["movie_id"]

	err := services.WatchMovie(mc.DB, movieID)
	var response models.Response
	if err != nil {
		response = models.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to increment watch num of movie with ID " + movieID,
			Data:       err,
		}
	} else {
		response = models.Response{
			StatusCode: http.StatusOK,
			Message:    "Success increment watch num of movie with ID " + movieID,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

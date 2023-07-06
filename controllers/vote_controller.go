package controllers

import (
	"encoding/json"
	"net/http"
	"rocketin-movie/models/dto"
	"rocketin-movie/models/extra"
	"rocketin-movie/services"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type VoteController struct {
	DB *gorm.DB
}

func NewVoteController(db *gorm.DB) *VoteController {
	return &VoteController{DB: db}
}

func (vc *VoteController) UpvoteMovie(w http.ResponseWriter, r *http.Request) {
	var dto dto.VoteDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	movieID := mux.Vars(r)["movie_id"]

	movie, err := services.UpvoteMovie(vc.DB, movieID, dto)
	var response extra.Response
	if err != nil {
		response = extra.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Failed to upvote movie with ID " + movieID,
			Data:       err.Error(),
		}
	} else {
		response = extra.Response{
			StatusCode: http.StatusOK,
			Message:    "Success upvote movie with ID " + movieID,
			Data:       movie,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (vc *VoteController) DownvoteMovie(w http.ResponseWriter, r *http.Request) {
	var dto dto.VoteDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	movieID := mux.Vars(r)["movie_id"]

	movie, err := services.DownVoteMovie(vc.DB, movieID, dto)
	var response extra.Response
	if err != nil {
		response = extra.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Failed to downvote movie with ID " + movieID,
			Data:       err.Error(),
		}
	} else {
		response = extra.Response{
			StatusCode: http.StatusOK,
			Message:    "Success downvote movie with ID " + movieID,
			Data:       movie,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

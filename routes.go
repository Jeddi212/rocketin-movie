package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"rocketin-movie/controllers"
)

func RegisterRoutes(db *gorm.DB, e *mux.Router) {

	movieController := controllers.NewMovieController(db)

	e.HandleFunc("/movies", movieController.GetMovie).Methods(http.MethodGet)
	e.HandleFunc("/create-movies", movieController.CreateMovie).Methods(http.MethodPost)
	e.HandleFunc("/movies/{movie_id}", movieController.UpdateMovie).Methods(http.MethodPut)
	e.HandleFunc("/watch/{movie_id}", movieController.WatchMovie).Methods(http.MethodPut)
}

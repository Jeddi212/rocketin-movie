package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"rocketin-movie/controllers"
)

func RegisterRoutes(db *gorm.DB, e *mux.Router) {

	movieController := controllers.NewMovieController(db)
	genreController := controllers.NewGenreController(db)
	mostController := controllers.NewMostController(db)
	voteController := controllers.NewVoteController(db)
	userController := controllers.NewUserController(db)

	e.HandleFunc("/movies-all", movieController.GetAllMovie).Methods(http.MethodGet)
	e.HandleFunc("/movies", movieController.GetMovie).Methods(http.MethodGet)
	e.HandleFunc("/create-movie", movieController.CreateMovie).Methods(http.MethodPost)
	e.HandleFunc("/movie/{movie_id}", movieController.UpdateMovie).Methods(http.MethodPut)
	e.HandleFunc("/watch/{movie_id}", movieController.WatchMovie).Methods(http.MethodPut)

	e.HandleFunc("/create-genre", genreController.CreateGenre).Methods(http.MethodPost)

	e.HandleFunc("/most-viewed", mostController.GetMostViewed).Methods(http.MethodGet)

	e.HandleFunc("/upvote/{movie_id}", voteController.UpvoteMovie).Methods(http.MethodPost)

	e.HandleFunc("/register", userController.Register).Methods(http.MethodPost)
}

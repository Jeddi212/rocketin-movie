package main

import (
	"log"
	"net/http"

	"fmt"

	"rocketin-movie/database"
	"rocketin-movie/routes"

	"github.com/gorilla/mux"
)

func main() {
	db := database.GetDBConnection()
	database.Migrate(db)

	r := mux.NewRouter()
	routes.RegisterRoutes(db, r)

	http.Handle("/", r)
	fmt.Println("Connected to port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

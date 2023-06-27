package main

import (
	"log"
	"net/http"

	"fmt"

	"github.com/gorilla/mux"
)

func main() {
	db := GetDBConnection()

	Migrate(db)

	r := mux.NewRouter()
	RegisterRoutes(db, r)

	http.Handle("/", r)
	fmt.Println("Connected to port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

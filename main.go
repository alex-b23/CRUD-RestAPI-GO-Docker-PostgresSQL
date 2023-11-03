package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// the structure of the player filesystem
type Player struct {
	ID 			   int 		`json:"id"`
	PlayerUsername string 	`json:"player_username"`
	PlayerPassword string 	`json:"player_password"`
	PlayerWins 	   int 		`json:"player_wins"`
	PlayerLoses    int 		`json:"player_loses"`
	PlayerTotalGame int 	`json:"player_total_game"`
}

func main() {
	// Connecting to SQL Database
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL")) // TO DO -> Change URL once Docker Setup

	if err != nil {
		log.Fatal(err)
	}
	// Close DB Connection
	defer db.Close()

	// Creating Router Requests
	route := mux.NewRouter()
	route.HandleFunc("/players", getPlayer(db)).Methods("GET") // getPlayer -> used to get players from database
	route.HandleFunc("/players/{id}", getPlayerByID(db)).Methods("GET") // getPlayerByID -> used to get player by ID
	route.HandleFunc("/players", createPlayer(db)).Methods("POST") // createPlayer -> add a new player to the database
	route.HandleFunc("/players/{id}", updatePlayer(db)).Methods("PUT") // updateUser -> update player information
	route.HandleFunc("/players/{id}", deletePlayer(db)).Methods("DELETE") // deletePlayer -> delete player from database

	// used to start the server
	log.Fatal(http.ListenAndServe(":3000", jsonMiddleware(route)))
}

// This Function will send our data as JSON to the server
func jsonMiddleware (next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
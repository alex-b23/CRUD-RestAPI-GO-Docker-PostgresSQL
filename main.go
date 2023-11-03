package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
    db, err := sql.Open("postgres", os.Getenv("DATABASE_URL")) // TO DO -> Connect Database Once Setup
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    route := mux.NewRouter()
    route.HandleFunc("/players", getPlayer(db)).Methods("GET") // -> method used to get all players
    route.HandleFunc("/players/{id}", getPlayerByID(db)).Methods("GET") // -> method used to get player by id
    route.HandleFunc("/players", createPlayer(db)).Methods("POST") // -> method used to create new player
    route.HandleFunc("/players/{id}", updatePlayer(db)).Methods("PUT") // -> method used to update player by id
    route.HandleFunc("/players/{id}", deletePlayer(db)).Methods("DELETE") // -> method used to delete player by id

    log.Fatal(http.ListenAndServe(":3000", jsonMiddleware(route)))
}
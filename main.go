package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
    db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    route := mux.NewRouter()
    route.HandleFunc("/players", getPlayer(db)).Methods("GET")
    route.HandleFunc("/players/{id}", getPlayerByID(db)).Methods("GET")
    route.HandleFunc("/players", createPlayer(db)).Methods("POST")
    route.HandleFunc("/players/{id}", updatePlayer(db)).Methods("PUT")
    route.HandleFunc("/players/{id}", deletePlayer(db)).Methods("DELETE")

    log.Fatal(http.ListenAndServe(":3000", jsonMiddleware(route)))
}
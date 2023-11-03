package main

import (
	"database/sql"
	"net/http"
)

type Player struct {
    ID             int    `json:"id"`
    PlayerUsername string `json:"player_username"`
    PlayerPassword string `json:"player_password"`
    PlayerWins     int    `json:"player_wins"`
    PlayerLoses    int    `json:"player_loses"`
    PlayerTotalGame int   `json:"player_total_game"`
}

func getPlayer(db *sql.DB) http.HandlerFunc {
    // implementation here
}

func getPlayerByID(db *sql.DB) http.HandlerFunc {
    // implementation here
}

func createPlayer(db *sql.DB) http.HandlerFunc {
    // implementation here
}

func updatePlayer(db *sql.DB) http.HandlerFunc {
    // implementation here
}

func deletePlayer(db *sql.DB) http.HandlerFunc {
    // implementation here
}
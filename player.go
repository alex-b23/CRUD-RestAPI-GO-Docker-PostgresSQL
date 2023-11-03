package main

import (
	"database/sql"
	"encoding/json"
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

// this function it's used to get all players from the database
func getPlayer(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        rows, err := db.Query("SELECT * FROM players")
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        defer rows.Close()

        players := []Player{}
        for rows.Next() {
            var u Player
            if err := rows.Scan(&u.ID, &u.PlayerUsername, &u.PlayerPassword, &u.PlayerWins, &u.PlayerLoses, &u.PlayerTotalGame); err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            players = append(players, u)
        }
        if err := rows.Err(); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        if err := json.NewEncoder(w).Encode(players); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    }
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
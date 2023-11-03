package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
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

// this function will get player by the specified id
func getPlayerByID(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        id, ok := vars["id"]
        if !ok {
            http.Error(w, "Missing URL parameter id", http.StatusBadRequest)
            return
        }

        var u Player
        err := db.QueryRow("SELECT * FROM players WHERE id = $1", id).Scan(&u.ID, &u.PlayerUsername, &u.PlayerPassword, &u.PlayerWins, &u.PlayerLoses, &u.PlayerTotalGame)
        if err == sql.ErrNoRows {
            http.Error(w, "No player with that ID", http.StatusNotFound)
            return
        } else if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        if err := json.NewEncoder(w).Encode(u); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    }
}

// This function creates a new player to the database
func createPlayer(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var u Player
        if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        sqlStatement := `
        INSERT INTO players (player_username, player_password, player_wins, player_loses, player_total_game)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id`
        id := 0
        err := db.QueryRow(sqlStatement, u.PlayerUsername, u.PlayerPassword, u.PlayerWins, u.PlayerLoses, u.PlayerTotalGame).Scan(&id)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusCreated)
        response := map[string]int{"id": id}
        if err := json.NewEncoder(w).Encode(response); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    }
}

// this function updates player
func updatePlayer(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        id, ok := vars["id"]
        if !ok {
            http.Error(w, "Missing URL parameter id", http.StatusBadRequest)
            return
        }

        var u Player
        if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        sqlStatement := `
        UPDATE players 
        SET player_username = $1, player_password = $2, player_wins = $3, player_loses = $4, player_total_game = $5
        WHERE id = $6
        RETURNING id, player_username, player_password, player_wins, player_loses, player_total_game`
        err := db.QueryRow(sqlStatement, u.PlayerUsername, u.PlayerPassword, u.PlayerWins, u.PlayerLoses, u.PlayerTotalGame, id).Scan(&u.ID, &u.PlayerUsername, &u.PlayerPassword, &u.PlayerWins, &u.PlayerLoses, &u.PlayerTotalGame)
        if err == sql.ErrNoRows {
            http.Error(w, "No player with that ID", http.StatusNotFound)
            return
        } else if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        if err := json.NewEncoder(w).Encode(u); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    }
}

// this function deletes a player from the database given the ID
func deletePlayer(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        id, ok := vars["id"]
        if !ok {
            http.Error(w, "Missing URL parameter id", http.StatusBadRequest)
            return
        }

        sqlStatement := `DELETE FROM players WHERE id = $1`
        _, err := db.Exec(sqlStatement, id)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusNoContent)
    }
}
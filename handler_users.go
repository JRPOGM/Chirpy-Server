package main

import (
	"encoding/json"
	"net/http"
	"log"
	"time"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
}

func (cfg *apiConfig) handlerUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	type parameter struct {
		Password string `json:"password"`
		Email string `json:"email"`
	}
	decoder := json.NewDecoder(r.Body)
	param := parameter{}
	err := decoder.Decode(&param)
	if err != nil {
		log.Printf("Could not decode the user: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	hashPass, err := HashPassword(param.Password)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't hash the password", err)
	}
	dbUser, err := cfg.db.CreateUser(ctx, param.Email)
	respUser := User{
		ID: dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Email: dbUser.Email,
	}
	respondWithJson(w, http.StatusCreated, respUser)
}
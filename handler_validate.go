package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"github.com/JRPOGM/Chirpy-Server/internal/database"
)

type Chirps struct {
	ID        uuid.UUID `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    Body   string    `json:"content"`
    UserID    uuid.UUID `json:"user_id"`
}

func (cfg *apiConfig) handlerChirps(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body 	string 		`json:"body"`
		UserID 	uuid.UUID 	`json:"user_id"`
	}
	type returnValues struct {
		CleanBody string `json:"cleaned_body"`
	}
	decoder := json.NewDecoder(r.Body)
	param := parameters{}
	err := decoder.Decode(&param)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}
	const maxChirpString = 140
	if len(param.Body) > maxChirpString {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
		return
	}
	badWords := map[string]struct{}{
		"kerfuffle": {},
		"sharbert": {},
		"fornax": {},
	}
	clean, err := getCleanBody(param.Body, badWords)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not clean the text", err)
		return
	}
	chirp, err := cfg.db.CreateChirps(r.Context(), database.CreateChirpsParams{
		Body: clean,
		UserID: param.UserID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't create chirps", err)
		return
	}
	respChirps := Chirps{
		ID:			chirp.ID,
		CreatedAt:	chirp.CreatedAt,
		UpdatedAt:	chirp.UpdatedAt,
		Body:		chirp.Body,
		UserID:		chirp.UserID,
	}
	respondWithJson(w, http.StatusCreated, respChirps)
}

func getCleanBody(body string, badWords map[string]struct{}) string {
	words := strings.Split(body, " ")
	for i, word := range words {
		lowWord := strings.ToLower(word)
		if _, ok := badWords[lowWord]; ok {
			words[i] = "****"
		}
	}
	clean := strings.Join(words, " ")
	return clean
}
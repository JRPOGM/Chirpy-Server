package main

import (
	"net/http"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerGetSingleChirp(w http.ResponseWriter, r *http.Request) {
	chirpIDString := r.PathValue("chirpID")
	chirpID, err := uuid.Parse(chirpIDString)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Chirp ID", err)
		return
	}
	dbChirp, err := cfg.db.GetChirp(r.Context(), chirpID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Chirp could not be located", err)
		return
	}
	chirp := Chirp{
		ID:			dbChirp.ID,
		CreatedAt:	dbChirp.CreatedAt,
		UpdatedAt:	dbChirp.UpdatedAt,
		Body:		dbChirp.Body,
		UserID:		dbChirp.UserID,
	}
	respondWithJson(w, http.StatusOK, chirp)
}
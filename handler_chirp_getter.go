package main

import (
	"net/http"
)

func (cfg *apiConfig) handlerGetChirps(w http.ResponseWriter, r *http.Request) {
	chirp, err := cfg.db.GetChirps(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't find the chirps", err)
		return
	}
	respondWithJson(w, http.StatusOK, chirp)
}
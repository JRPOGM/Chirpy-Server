package main

import (
	"net/http"
)

type Chirps struct {
	ID        uuid.UUID `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    Body	  string    `json:"body"`
    UserID    uuid.UUID `json:"user_id"`
}

func (cfg *apiConfig) handlerGetChirps(w http.ResponseWriter, r *http.Request) {
	dbChirp, err := cfg.db.GetChirps(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't find the chirps", err)
		return
	}
	chirps := []Chirp{}
	for _, chirp := range dbChirp {
		chirps = append(chirps, Chirp{
			ID: chirp.ID,
			CreatedAt: chirp.CreatedAt,
			UpdatedAt: chirp.UpdatedAt,
			Body: chirp.Body,
			UserID: chirp.UserID,
		})
	}
	respondWithJson(w, http.StatusOK, chirps)
}
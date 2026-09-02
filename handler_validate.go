package main

import (
	"encoding/json"
	"net/http"
)

func handlerValidateChirp(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}
	type returnValues struct {
		Valid bool `json:"valid"`
	}
	decoder := json.NewDecoder(r.Body)
	param := parameters{}
	err := decoder.Decode(&param)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}
	const maxChirpString = 140
	if len(params.Body) > maxChirpString {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
		return
	}
	respondWithJson(w, http.StatusOK, returnValues{
		Valid: true,
	})
}
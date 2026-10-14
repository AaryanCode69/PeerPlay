package main

import (
	"encoding/json"
	"net/http"
)

func handleRequestBody(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	params := parameters{}

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't decode parameters", err)
		return
	}

	respondWithJSON(w, http.StatusOK, params)
}

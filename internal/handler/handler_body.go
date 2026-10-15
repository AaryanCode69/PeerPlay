// Package handler handles request routing
package handler

import (
	"encoding/json"
	"net/http"

	"PeerPlay/internal/httputil"
)

func HandleRequestBody(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	params := parameters{}

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		httputil.RespondWithError(w, http.StatusBadRequest, "Couldn't decode parameters", err)
		return
	}

	httputil.RespondWithJSON(w, http.StatusOK, params)
}

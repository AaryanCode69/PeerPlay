// Package httputil provides util functions for code reusability and readability
package httputil

import (
	"encoding/json"
	"log"
	"net/http"
)

// RespondWithJSON turns any Go value into JSON and sends it
// with the given status code.
func RespondWithJSON(w http.ResponseWriter, code int, payload any) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

// RespondWithError sends {"error": "msg"} to the client.
// If err is not nil, we log it for ourselves but never show it to the client.
func RespondWithError(w http.ResponseWriter, code int, msg string, err error) {
	if err != nil {
		log.Println(err)
	}
	if code >= 500 {
		log.Printf("Responding with 5XX error: %s", msg)
	}

	type errorResponse struct {
		Error string `json:"error"`
	}
	RespondWithJSON(w, code, errorResponse{Error: msg})
}

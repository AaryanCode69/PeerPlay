package handler

import (
	"net/http"

	"PeerPlay/internal/httputil"
)

func HandlerHealth(w http.ResponseWriter, r *http.Request) {
	httputil.RespondWithJSON(w, 200, struct {
		Status string `json:"status"`
	}{
		Status: "ok",
	})
}

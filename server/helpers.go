package server

import (
	"encoding/json"
	"gitlab.com/stz184/gopher-translator/models"
	"log"
	"net/http"
)

func ToJSON(w http.ResponseWriter, r interface{}) {
	w.Header().Set("Content-type", "application/json")
	err := json.NewEncoder(w).Encode(r)
	if err != nil {
		log.Printf("Failed to serialize the response: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func RespondWithError(w http.ResponseWriter, error string) {
	var errorResponse = models.ErrorResponse{Error: error}
	ToJSON(w, errorResponse)
}

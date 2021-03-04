package server

import (
	"encoding/json"
	"github.com/stz184/gopher-translator/models"
	"log"
	"net/http"
)

func encodeOutput(w http.ResponseWriter, r interface{}, statusCode int) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(r)
	if err != nil {
		log.Printf("Failed to serialize the response: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func ToJSON(w http.ResponseWriter, r interface{}) {
	encodeOutput(w, r, http.StatusOK)
}

func RespondWithError(w http.ResponseWriter, error string) {
	var errorResponse = models.ErrorResponse{Error: error}
	encodeOutput(w, errorResponse, http.StatusBadRequest)
}

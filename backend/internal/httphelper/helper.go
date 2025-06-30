package httphelper

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}

func WriteJSONError(w http.ResponseWriter, status int, message string) {
	log.Printf("Error %d: %s", status, message) // log côté serveur

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	resp := ErrorResponse{
		Error: message,
		Code:  status,
	}
	json.NewEncoder(w).Encode(resp)
}

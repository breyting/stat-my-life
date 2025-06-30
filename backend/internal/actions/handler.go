package actions

import (
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v5"
)

type Handler struct {
	DB *pgx.Conn
}

// GET /actions
func (h *Handler) GetActions(w http.ResponseWriter, r *http.Request) {
	actions, err := GetAll(r.Context(), h.DB)
	if err != nil {
		http.Error(w, "Failed to fetch actions", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(actions)
}

// POST /actions
func (h *Handler) CreateAction(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	action, err := Insert(r.Context(), h.DB, body.Name)
	if err != nil {
		http.Error(w, "Failed to insert action", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(action)
}

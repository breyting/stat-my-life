package actions

import (
	"encoding/json"
	"net/http"
	"time"

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

// GET /actions/{id}/events
func (h *Handler) GetOccurrences(w http.ResponseWriter, r *http.Request, actionID int) {
	occurrences, err := GetOccurrencesByAction(r.Context(), h.DB, actionID)
	if err != nil {
		http.Error(w, "Failed to fetch occurrences", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(occurrences)
}

// POST /actions/{id}/events
func (h *Handler) CreateOccurrence(w http.ResponseWriter, r *http.Request, actionID int) {
	var body struct {
		Timestamp *time.Time `json:"timestamp,omitempty"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Par défaut, timestamp = maintenant
	ts := time.Now().UTC()
	if body.Timestamp != nil {
		ts = *body.Timestamp
	}

	occurrence, err := InsertOccurrence(r.Context(), h.DB, actionID, ts)
	if err != nil {
		http.Error(w, "Failed to insert occurrence", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(occurrence)
}

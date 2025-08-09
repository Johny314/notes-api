package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/jeka314/notes-api/models"
	"github.com/jeka314/notes-api/storage"
)

// GetNotes godoc
// @Summary Get all notes
// @Description Returns a list of all notes
// @Tags notes
// @Produce json
// @Success 200 {array} models.Note
// @Router /notes [get]
func GetNotes(w http.ResponseWriter, r *http.Request) {
	var notes []models.Note
	storage.DB.Find(&notes)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)
}

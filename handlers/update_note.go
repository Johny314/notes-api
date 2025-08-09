package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jeka314/notes-api/models"
	"github.com/jeka314/notes-api/storage"
)

// UpdateNote godoc
// @Summary Update a note by ID
// @Description Updates title and content of a note
// @Tags notes
// @Accept json
// @Produce json
// @Param id path int true "Note ID"
// @Param note body models.Note true "Updated note object"
// @Success 200 {object} models.Note
// @Failure 400 {string} string "Invalid input"
// @Failure 404 {string} string "Note not found"
// @Router /notes/{id} [put]
func UpdateNote(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var note models.Note
	if err := storage.DB.First(&note, id).Error; err != nil {
		http.Error(w, "Note not found", http.StatusNotFound)
		return
	}

	var updated models.Note
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	note.Title = updated.Title
	note.Content = updated.Content
	storage.DB.Save(&note)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

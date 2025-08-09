package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/jeka314/notes-api/models"
	"github.com/jeka314/notes-api/storage"
)

// CreateNote godoc
// @Summary Create a new note
// @Description Adds a new note to the database
// @Tags notes
// @Accept json
// @Produce json
// @Param note body models.Note true "Note object"
// @Success 200 {object} models.Note
// @Failure 400 {string} string
// @Router /notes [post]
func CreateNote(w http.ResponseWriter, r *http.Request) {
	var note models.Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	storage.DB.Create(&note)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

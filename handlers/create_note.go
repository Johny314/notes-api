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
// @Failure 400 {object} map[string]string
// @Router /notes [post]
func CreateNote(w http.ResponseWriter, r *http.Request) {
	var note models.Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		HandleError(w, err, "Failed to decode note", http.StatusBadRequest)
		return
	}

	if err := validate.Struct(note); err != nil {
		HandleError(w, err, "Validation failed: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := storage.DB.Create(&note).Error; err != nil {
		HandleError(w, err, "Failed to create note", http.StatusInternalServerError)
		return
	}

	SendSuccess(w, note)
}

package handlers

import (
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
func GetNotes(w http.ResponseWriter, _ *http.Request) {
	var notes []models.Note
	if err := storage.DB.Find(&notes).Error; err != nil {
		HandleError(w, err, "Failed to get notes", http.StatusInternalServerError)
		return
	}

	SendSuccess(w, notes)
}

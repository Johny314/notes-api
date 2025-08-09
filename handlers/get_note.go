package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jeka314/notes-api/models"
	"github.com/jeka314/notes-api/storage"
)

// GetNote godoc
// @Summary Get a note by ID
// @Description Returns a single note by its ID
// @Tags notes
// @Produce json
// @Param id path int true "Note ID"
// @Success 200 {object} models.Note
// @Failure 404 {object} map[string]string
// @Router /notes/{id} [get]
func GetNote(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var note models.Note
	if err := storage.DB.First(&note, id).Error; err != nil {
		HandleError(w, err, "Note not found", http.StatusNotFound)
		return
	}

	SendSuccess(w, note)
}

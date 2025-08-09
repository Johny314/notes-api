package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jeka314/notes-api/models"
	"github.com/jeka314/notes-api/storage"
)

// DeleteNote godoc
// @Summary Delete a note by ID
// @Description Deletes a note from the database
// @Tags notes
// @Param id path int true "Note ID"
// @Success 204 "No Content"
// @Failure 404 {object} map[string]string
// @Router /notes/{id} [delete]
func DeleteNote(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if err := storage.DB.Delete(&models.Note{}, id).Error; err != nil {
		HandleError(w, err, "Note not found", http.StatusNotFound)
		return
	}

	SendNoContent(w)
}

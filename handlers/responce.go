package handlers

import (
	"encoding/json"
	"net/http"
)

// SuccessResponse — стандартный формат успешного ответа
type SuccessResponse struct {
	Data interface{} `json:"data"`
}

// SendSuccess отправляет JSON-ответ с данными и статусом 200
func SendSuccess(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	resp := SuccessResponse{Data: data}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func SendNoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

package handlers

import (
	"net/http"

	"github.com/jeka314/notes-api/logger"
	"go.uber.org/zap"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

// HandleError логирует ошибку и отправляет JSON-ответ с сообщением об ошибке
func HandleError(w http.ResponseWriter, err error, message string, statusCode int) {
	logger.Log.Error(message, zap.Error(err))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	resp := ErrorResponse{Error: message}
	SendSuccess(w, resp)
}

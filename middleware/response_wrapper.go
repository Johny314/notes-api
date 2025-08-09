package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/jeka314/notes-api/logger"
)

type ResponseWriter struct {
	http.ResponseWriter
	statusCode int
	body       []byte
}

func NewResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}
}

func (rw *ResponseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *ResponseWriter) Write(b []byte) (int, error) {
	rw.body = append(rw.body, b...)
	return rw.ResponseWriter.Write(b)
}

type SuccessResponse struct {
	Data interface{} `json:"data"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func ResponseWrapper(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw := NewResponseWriter(w)
		next.ServeHTTP(rw, r)

		if rw.statusCode >= 200 && rw.statusCode < 300 {
			var originalData interface{}
			if err := json.Unmarshal(rw.body, &originalData); err != nil {
				// Если не JSON, просто вывести как есть
				logger.Log.Warn("Response is not valid JSON, skipping wrapper")
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(rw.statusCode)
			resp := SuccessResponse{Data: originalData}
			_ = json.NewEncoder(w).Encode(resp)
		} else if rw.statusCode >= 400 {
			// Ошибка — ожидаем что тело уже содержит {"error": "..."} или plain text
			// Если не JSON, преобразуем plain text в json
			var jsonCheck interface{}
			if err := json.Unmarshal(rw.body, &jsonCheck); err != nil {
				// Тело не JSON — оборачиваем в ErrorResponse с plain текстом
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(rw.statusCode)
				resp := ErrorResponse{Error: string(rw.body)}
				_ = json.NewEncoder(w).Encode(resp)
			}
		}
	})
}

package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jeka314/notes-api/handlers"
	"github.com/jeka314/notes-api/logger"
	"github.com/jeka314/notes-api/middleware"
	"github.com/jeka314/notes-api/storage"
	"go.uber.org/zap"

	_ "github.com/jeka314/notes-api/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	if err := logger.Init(); err != nil {
		panic(err)
	}
	defer logger.Sync()

	storage.InitDB()

	r := mux.NewRouter()

	r.Use(middleware.ResponseWrapper)
	r.Use(loggingMiddleware)

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	r.HandleFunc("/notes", handlers.CreateNote).Methods("POST")
	r.HandleFunc("/notes", handlers.GetNotes).Methods("GET")
	r.HandleFunc("/notes/{id}", handlers.GetNote).Methods("GET")
	r.HandleFunc("/notes/{id}", handlers.UpdateNote).Methods("PUT")
	r.HandleFunc("/notes/{id}", handlers.DeleteNote).Methods("DELETE")

	logger.Log.Info("Server started at :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		logger.Log.Fatal("Server failed to start", zap.Error(err))
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Log.Info("HTTP request",
			zap.String("method", r.Method),
			zap.String("url", r.RequestURI),
			zap.String("remote_addr", r.RemoteAddr),
		)
		next.ServeHTTP(w, r)
	})
}

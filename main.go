package main

import (
	"fmt"
	"net/http"

	_ "github.com/jeka314/notes-api/docs" // docs сгенерирует swag
	"github.com/jeka314/notes-api/handlers"
	"github.com/jeka314/notes-api/storage"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Notes API
// @version 1.0
// @description Simple REST API for managing notes.
// @BasePath /

// @host localhost:8080
func main() {
	// Инициализация базы данных
	storage.InitDB()

	r := mux.NewRouter()

	// Swagger UI
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Роуты API
	r.HandleFunc("/notes", handlers.CreateNote).Methods("POST")
	r.HandleFunc("/notes", handlers.GetNotes).Methods("GET")
	r.HandleFunc("/notes/{id}", handlers.GetNote).Methods("GET")
	r.HandleFunc("/notes/{id}", handlers.UpdateNote).Methods("PUT")
	r.HandleFunc("/notes/{id}", handlers.DeleteNote).Methods("DELETE")

	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", r)
}

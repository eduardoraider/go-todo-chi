package main

import (
	"fmt"
	"github.com/eduardoraider/go-todo-chi/internal/handlers"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	router := chi.NewRouter()

	router.Get("/todos", handlers.GetTodosHandler)
	router.Get("/todos/{id}", handlers.GetChiTodoByID)
	router.Post("/todos", handlers.CreateTodoHandler)
	router.Put("/todos/{id}", handlers.UpdateTodoHandler)
	router.Delete("/todos/{id}", handlers.DeleteTodoHandler)

	// Start server
	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", router)
}

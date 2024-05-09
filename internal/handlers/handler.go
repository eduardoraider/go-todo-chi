package handlers

import (
	"encoding/json"
	"github.com/eduardoraider/go-todo-chi/internal/entity"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func GetTodosHandler(w http.ResponseWriter, r *http.Request) {
	todos := entity.GetTodos()

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func GetChiTodoByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	todo := entity.GetTodoById(id)

	if err := json.NewEncoder(w).Encode(todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}

func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todo entity.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	entity.AddTodo(todo)

	w.WriteHeader(http.StatusCreated)
}

func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var todo entity.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	entity.UpdateTodo(id, todo)

	w.WriteHeader(http.StatusAccepted)
}

func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	entity.DeleteTodo(id)
	http.Error(w, "Todo not found", http.StatusNoContent)
}

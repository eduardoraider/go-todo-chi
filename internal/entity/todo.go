package entity

type Todo struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

// Slice to hold todo items
var todos []Todo

func GetTodos() []Todo {
	return todos
}

func GetTodoById(id string) *Todo {
	for _, todo := range todos {
		if todo.ID == id {
			return &todo
		}
	}
	return nil
}

func AddTodo(todo Todo) {
	todos = append(todos, todo)
}

func UpdateTodo(id string, updatedTodo Todo) {
	for i, todo := range todos {
		if todo.ID == id {
			todos[i] = updatedTodo
			break
		}
	}
}

func DeleteTodo(id string) {
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return
		}
	}
}

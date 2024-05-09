### Go TODO with Chi

This project implements a TODO API server in Go using Chi, a lightweight and flexible router for HTTP services.

#### Installation

To get started, install Chi using the following command:

```
go get -u github.com/go-chi/chi/v5
```

#### Usage

Run the API server with the following command:

```
go run cmd/api/main.go
```

#### API Endpoints

- **GET /todos**: Retrieve all todo items.
- **GET /todos/{id}**: Retrieve a todo item by its ID.
- **POST /todos**: Add a new todo item.
- **PUT /todos/{id}**: Update an existing todo item.
- **DELETE /todos/{id}**: Delete an existing todo item.

#### Using CLI (curl)

- **Get all todos**:
```
curl http://localhost:8080/todos
```

- **Get a todo by ID**:
```
curl http://localhost:8080/todos/1
```

- **Add a new todo**:
```
curl -X POST -H "Content-Type: application/json" -d '{"id": "1", "title": "New Todo", "done": false}' http://localhost:8080/todos
```

- **Update a todo**:
```
curl -X PUT -H "Content-Type: application/json" -d '{"id": "1", "title": "Updated Todo", "done": true}' http://localhost:8080/todos/2
```

- **Delete a todo**:
```
curl -X DELETE http://localhost:8080/todos/1
```

#### Using Postman

- **Get all todos**:
    - Method: GET
    - URL: http://localhost:8080/todos

- **Get a todo by ID**:
    - Method: GET
    - URL: http://localhost:8080/todos/1 (Replace 1 with the ID of the todo)

- **Add a new todo**:
    - Method: POST
    - URL: http://localhost:8080/todos
    - Headers:
        - Key: Content-Type
        - Value: application/json
    - Body (raw JSON):
      ```json
      {
          "id": "1",
          "title": "New Todo",
          "done": false
      }
      ```

- **Update a todo**:
    - Method: PUT
    - URL: http://localhost:8080/todos/1 (Replace 1 with the ID of the todo)
    - Headers:
        - Key: Content-Type
        - Value: application/json
    - Body (raw JSON):
      ```json
      {
          "id": "1",
          "title": "Updated Todo",
          "done": true
      }
      ```

- **Delete a todo**:
    - Method: DELETE
    - URL: http://localhost:8080/todos/1 (Replace 1 with the ID of the todo)


## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

by **Eduardo Raider** - Software Engineer
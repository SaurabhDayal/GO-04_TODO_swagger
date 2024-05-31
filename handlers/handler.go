package handlers

import (
	"04_todo_swagger/database/dbHelper"
	"04_todo_swagger/models"
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// ReadTodo godoc
// @Summary Read a single todo
// @Description Get a todo by ID
// @ID read-todo
// @Param todoId path int true "Todo ID"
// @Produce json
// @Success 200 {object} Todo
// @Failure 204 {string} string "No content"
// @Failure 500 {string} string "Internal server error"
// @Router /todo/{todoId} [get]
func ReadTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	strId := vars["todoId"]
	id, _ := strconv.Atoi(strId)
	todo, err := dbHelper.FindTodoById(id)
	if err != nil && err != sql.ErrNoRows {
		w.WriteHeader(http.StatusInternalServerError)
	} else if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(todo)
	}
}

// ReadAllTodo godoc
// @Summary Read all todos
// @Description Get all todos
// @ID read-all-todo
// @Produce json
// @Success 200 {array} Todo
// @Failure 204 {string} string "No content"
// @Router /todo [get]
func ReadAllTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	todos, err := dbHelper.FindAllTodos()
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(todos)
	}
}

// CreateTodo godoc
// @Summary Create a todo
// @Description Create a new todo
// @ID create-todo
// @Accept json
// @Param userId path int true "User ID"
// @Param todo body Todo true "Todo object"
// @Success 200 {string} string "OK"
// @Failure 204 {string} string "No content"
// @Router /todo/{userId} [post]
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	strId := vars["userId"]
	id, _ := strconv.Atoi(strId)
	var todo models.Todo
	json.NewDecoder(r.Body).Decode(&todo)
	err := dbHelper.CreateNewTodo(&todo, id)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

// UpdateTodo godoc
// @Summary Update a todo
// @Description Update an existing todo
// @ID update-todo
// @Accept json
// @Param todoId path int true "Todo ID"
// @Param todo body Todo true "Updated todo object"
// @Success 200 {object} Todo
// @Failure 204 {string} string "No content"
// @Failure 500 {string} string "Internal server error"
// @Router /todo/{todoId} [put]
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	strId := vars["todoId"]
	id, _ := strconv.Atoi(strId)

	var todo models.Todo
	json.NewDecoder(r.Body).Decode(&todo)
	todo2, err := dbHelper.UpdateTodoById(todo, id)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusInternalServerError)
	} else if errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(todo2)
	}
}

// DeleteTodo godoc
// @Summary Delete a todo
// @Description Delete an existing todo
// @ID delete-todo
// @Param todoId path int true "Todo ID"
// @Success 200 {string} string "OK"
// @Failure 204 {string} string "No content"
// @Router /todo/{todoId} [delete]
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	strId := vars["todoId"]
	id, _ := strconv.Atoi(strId)

	err := dbHelper.DeleteTodoById(id)

	if err != nil {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

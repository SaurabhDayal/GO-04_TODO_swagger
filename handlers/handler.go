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

// ReadTask godoc
// @Summary Read a single task
// @Description Get a task by ID
// @ID read-task
// @Param todoId path int true "Task ID"
// @Produce json
// @Success 200 {object} Task
// @Failure 204 {string} string "No content"
// @Failure 500 {string} string "Internal server error"
// @Router /task/{taskId} [get]
func ReadTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	strId := vars["todoId"]
	id, _ := strconv.Atoi(strId)
	todo, err := dbHelper.FindTodoById(id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusInternalServerError)
	} else if errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(todo)
	}
}

// ReadAllTask godoc
// @Summary Read all todos
// @Description Get all todos
// @ID read-all-task
// @Produce json
// @Success 200 {array} Task
// @Failure 204 {string} string "No content"
// @Router /task [get]
func ReadAllTask(w http.ResponseWriter, r *http.Request) {
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
// @Summary Create a task
// @Description Create a new task
// @ID create-task
// @Accept json
// @Param userId path int true "User ID"
// @Param task body Task true "Task object"
// @Success 200 {string} string "OK"
// @Failure 204 {string} string "No content"
// @Router /task/{userId} [post]
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

// UpdateTask godoc
// @Summary Update a task
// @Description Update an existing task
// @ID update-task
// @Accept json
// @Param todoId path int true "Task ID"
// @Param task body Task true "Updated task object"
// @Success 200 {object} Task
// @Failure 204 {string} string "No content"
// @Failure 500 {string} string "Internal server error"
// @Router /task/{todoId} [put]
func UpdateTask(w http.ResponseWriter, r *http.Request) {
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

// DeleteTask godoc
// @Summary Delete a task
// @Description Delete an existing task
// @ID delete-task
// @Param todoId path int true "Task ID"
// @Success 200 {string} string "OK"
// @Failure 204 {string} string "No content"
// @Router /task/{taskId} [delete]
func DeleteTask(w http.ResponseWriter, r *http.Request) {
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

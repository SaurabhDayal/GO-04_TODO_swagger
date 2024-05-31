package handlers

import (
	"04_todo_swagger/database/dbHelper"
	"04_todo_swagger/models"
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// @Summary Get an item by ID
// @Description Get a single item by its ID
// @ID get-item-by-id
// @Produce json
// @Param id path int true "Item ID"
// @Success 200 {object} Item
// @Failure 400 {object} ErrorResponse
// @Router /items/{id} [get]
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

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	strId := vars["todoId"]
	id, _ := strconv.Atoi(strId)

	var todo models.Todo
	json.NewDecoder(r.Body).Decode(&todo)
	todo2, err := dbHelper.UpdateTodoById(todo, id)

	if err != nil && err != sql.ErrNoRows {
		w.WriteHeader(http.StatusInternalServerError)
	} else if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(todo2)
	}
}

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

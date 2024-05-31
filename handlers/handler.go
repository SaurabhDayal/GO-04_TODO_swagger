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

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var usr models.UserClient
	err := json.NewDecoder(r.Body).Decode(&usr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		user, err := dbHelper.RegisterUser(&usr)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	}
}

func VerifyUserMidd(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqToken := r.Header.Get("Authorization")
		vars := mux.Vars(r)
		strId := vars["userId"]
		id, _ := strconv.Atoi(strId)
		checkUser := dbHelper.CheckUser(reqToken, id)
		if checkUser == true {
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var usr models.UserClient
	err := json.NewDecoder(r.Body).Decode(&usr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	pwd := usr.Password
	var auth *models.UserSession
	vars := mux.Vars(r)
	strId := vars["userId"]
	id, _ := strconv.Atoi(strId)
	auth, err = dbHelper.LoginUser(id, pwd)

	if err != nil {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(auth)
	}
}

func LogoutUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strId := vars["userId"]
	id, _ := strconv.Atoi(strId)
	err := dbHelper.LogoutUser(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

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

func ToggleTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	strId := vars["todoId"]
	id, _ := strconv.Atoi(strId)

	todo, err := dbHelper.ToggleTodoById(id)
	if err != nil && err != sql.ErrNoRows {
		w.WriteHeader(http.StatusInternalServerError)
	} else if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(todo)
	}
}

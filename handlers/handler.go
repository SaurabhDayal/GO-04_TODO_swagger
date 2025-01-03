package handlers

import (
	"04_todo_swagger/database/dbHelper"
	"04_todo_swagger/models"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi"
)

// ReadTask godoc
// @Summary Read a single task
// @Description Get a task by ID
// @ID read-task
// @Param taskId path string true "Task ID"
// @Produce json
// @Success 200 {object} models.Task
// @Failure 204 {string} string "No content"
// @Failure 500 {string} string "Internal server error"
// @Router /task/{taskId} [get]
func ReadTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := chi.URLParam(r, "taskId")
	task, err := dbHelper.FindTasksById(id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusInternalServerError)
	} else if errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(task)
	}
}

// ReadAllTask godoc
// @Summary Read all tasks
// @Description Get all tasks
// @ID read-all-task
// @Produce json
// @Success 200 {array} models.Task
// @Failure 204 {string} string "No content"
// @Router /task [get]
func ReadAllTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tasks, err := dbHelper.FindAllTasks()
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(tasks)
	}
}

// CreateTask godoc
// @Summary Create a task
// @Description Create a new task
// @ID create-task
// @Accept json
// @Param task body models.Task true "Task object"
// @Success 201 {string} string "Created"
// @Failure 500 {string} string "Internal server error"
// @Router /task [post]
func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)
	err := dbHelper.CreateNewTask(&task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

// UpdateTask godoc
// @Summary Update a task
// @Description Update an existing task
// @ID update-task
// @Accept json
// @Param taskId path string true "Task ID"
// @Param task body models.Task true "Updated task object"
// @Success 200 {object} models.Task
// @Failure 204 {string} string "No content"
// @Failure 500 {string} string "Internal server error"
// @Router /task/{taskId} [put]
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := chi.URLParam(r, "taskId")
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)
	response, err := dbHelper.UpdateTaskById(task, id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusInternalServerError)
	} else if errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

// DeleteTask godoc
// @Summary Delete a task
// @Description Delete an existing task
// @ID delete-task
// @Param taskId path string true "Task ID"
// @Success 200 {string} string "OK"
// @Failure 204 {string} string "No content"
// @Failure 500 {string} string "Internal server error"
// @Router /task/{taskId} [delete]
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := chi.URLParam(r, "taskId")
	err := dbHelper.DeleteTaskById(id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusInternalServerError)
	} else if errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

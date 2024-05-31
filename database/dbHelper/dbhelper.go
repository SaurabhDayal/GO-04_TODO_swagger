package dbHelper

import (
	"04_todo_swagger/database"
	"04_todo_swagger/models"
	"time"
)

func FindTodoById(id int) (models.TodoResponse, error) {
	SQL := `SELECT user_id, title, description, status, pending_at, created_at FROM todos WHERE id = $1 AND archived_at IS NULL`
	var todo models.TodoResponse
	err := database.Todo.Get(&todo, SQL, id)
	if err != nil {
		return models.TodoResponse{}, err
	}
	return todo, nil
}

func FindAllTodos() ([]models.TodoResponse, error) {
	SQL := `SELECT user_id, title, description, status, pending_at, created_at FROM todos WHERE archived_at IS NULL`
	todos := make([]models.TodoResponse, 0)
	err := database.Todo.Select(&todos, SQL)
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func CreateNewTodo(todo *models.Todo, id int) error {
	SQL := `INSERT INTO todos (user_id, title, description, pending_at) VALUES ($1,$2,$3,$4)`
	_, err := database.Todo.Exec(SQL, id, todo.Title, todo.Description, todo.PendingAt)
	if err != nil {
		return err
	}
	return nil
}

func UpdateTodoById(todo models.Todo, id int) (models.TodoResponse, error) {
	SQL := `UPDATE todos SET title=$1, description=$2, pending_at=$3 WHERE id=$4`
	_, err := database.Todo.Exec(SQL, todo.Title, todo.Description, todo.PendingAt, id)
	if err != nil {
		return models.TodoResponse{}, err
	}
	SQL2 := `SELECT user_id, title, description, status, pending_at, created_at FROM todos WHERE id = $1 AND archived_at IS NULL`
	var todo2 models.TodoResponse
	err2 := database.Todo.Get(&todo2, SQL2, id)
	if err2 != nil {
		return models.TodoResponse{}, err2
	}
	return todo2, nil
}

func DeleteTodoById(id int) error {
	SQL := `UPDATE todos SET archived_at =$1 WHERE id=$2 AND archived_at IS NULL`
	_, err := database.Todo.Exec(SQL, time.Now(), id)
	if err != nil {
		return err
	}
	return nil
}

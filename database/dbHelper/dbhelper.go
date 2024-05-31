package dbHelper

import (
	"04_todo_swagger/database"
	"04_todo_swagger/models"
	"time"
)

func FindTodoById(id int) (models.Task, error) {
	SQL := `SELECT id, 
       			   title, 
       			   description, 
       			   pending_at, 
       			   created_at 
			FROM tasks 
			WHERE id = $1 
			  AND archived_at IS NULL
`
	var todo models.Task
	err := database.Todo.Get(&todo, SQL, id)
	if err != nil {
		return models.Task{}, err
	}
	return todo, nil
}

func FindAllTodos() ([]models.Task, error) {
	SQL := `SELECT id, 
       			   title, 
       			   description, 
       			   pending_at, 
       			   created_at 
			FROM tasks 
			WHERE archived_at IS NULL
`
	tasks := make([]models.Task, 0)
	err := database.Todo.Select(&tasks, SQL)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func CreateNewTodo(todo *models.Task) error {
	SQL := `INSERT INTO tasks (
                   title, 
                   description, 
                   pending_at
                   ) 
			VALUES ($1,$2,$3)
`
	_, err := database.Todo.Exec(SQL, todo.Title, todo.Description, todo.PendingAt)
	if err != nil {
		return err
	}
	return nil
}

func UpdateTodoById(todo models.Task, id int) (models.Task, error) {
	SQL := `UPDATE tasks 
			SET title = $1, 
			    description= $2, 
			    pending_at= $3 
			WHERE id = $4
`
	_, err := database.Todo.Exec(SQL, todo.Title, todo.Description, todo.PendingAt, id)
	if err != nil {
		return models.Task{}, err
	}
	SQL2 := `SELECT id, 
       				title, 
       				description, 
       				pending_at, 
       				created_at 
			FROM tasks 
			WHERE id = $1 
			  AND archived_at IS NULL
`
	var todo2 models.Task
	err2 := database.Todo.Get(&todo2, SQL2, id)
	if err2 != nil {
		return models.Task{}, err2
	}
	return todo2, nil
}

func DeleteTodoById(id int) error {
	SQL := `UPDATE tasks SET archived_at =$1 WHERE id=$2 AND archived_at IS NULL`
	_, err := database.Todo.Exec(SQL, time.Now(), id)
	if err != nil {
		return err
	}
	return nil
}

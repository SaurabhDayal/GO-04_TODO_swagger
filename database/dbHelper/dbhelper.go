package dbHelper

import (
	"04_todo_swagger/database"
	"04_todo_swagger/models"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"time"
)

func RegisterUser(usr *models.UserClient) (*models.UserClient, error) {
	SQL1 := `INSERT INTO users (name, email, password) VALUES ($1,$2,$3)`
	_, err := database.Todo.Exec(SQL1, usr.Name, usr.Email, usr.Password)
	if err != nil {
		return nil, err
	}
	var userId int
	err = database.Todo.Get(&userId, "SELECT id FROM users WHERE email=$1", usr.Email)
	if err != nil {
		return nil, err
	}
	fmt.Println("user id", userId)
	SQL2 := `INSERT INTO user_roles (user_id, user_role) VALUES ($1, $2)`
	_, err = database.Todo.Exec(SQL2, userId, "user")
	if err != nil {
		return nil, err
	}
	user := *usr
	return &user, nil
}

func CheckUser(token string, id int) bool {
	SQL := `SELECT user_id FROM user_session WHERE session_token=$1`
	var checkId int
	err := database.Todo.Get(&checkId, SQL, token)
	if err != nil {
		return false
	}
	if checkId != id {
		return false
	}
	return true
}

func LoginUser(id int, pwd string) (*models.UserSession, error) {
	SQL2 := `SELECT password FROM users WHERE id=$1`
	var pswd string

	err := database.Todo.Get(&pswd, SQL2, id)
	if err != nil {
		return nil, err
	}

	if pswd == pwd {
		b := make([]byte, 6)
		if _, err := rand.Read(b); err != nil {
			return nil, err
		}
		b = []byte(hex.EncodeToString(b))
		SQL := `INSERT INTO user_session (user_id, session_token) VALUES ($1, $2)`
		_, err = database.Todo.Exec(SQL, id, b)
		if err != nil {
			return nil, err
		}
		var usrSn models.UserSession
		SQL2 := `SELECT user_id, session_token FROM user_session WHERE user_id=$1`
		err = database.Todo.Get(&usrSn, SQL2, id)
		if err != nil {
			return nil, err
		}
		return &usrSn, nil
	}
	return nil, errors.New("no match in password given and database password")
}

func LogoutUser(id int) error {
	SQL := `DELETE FROM user_session WHERE user_id = $1`
	_, err := database.Todo.Exec(SQL, id)
	if err != nil {
		return err
	}
	return nil
}

func FindTodoById(id int) (models.TodoClient, error) {
	SQL := `SELECT user_id, title, description, status, pending_at, created_at FROM todos WHERE id = $1 AND archived_at IS NULL`
	var todo models.TodoClient
	err := database.Todo.Get(&todo, SQL, id)
	if err != nil {
		return models.TodoClient{}, err
	}
	return todo, nil
}

func FindAllTodos() ([]models.TodoClient, error) {
	SQL := `SELECT user_id, title, description, status, pending_at, created_at FROM todos WHERE archived_at IS NULL`
	todos := make([]models.TodoClient, 0)
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

func UpdateTodoById(todo models.Todo, id int) (models.TodoClient, error) {
	SQL := `UPDATE todos SET title=$1, description=$2, pending_at=$3 WHERE id=$4`
	_, err := database.Todo.Exec(SQL, todo.Title, todo.Description, todo.PendingAt, id)
	if err != nil {
		return models.TodoClient{}, err
	}
	SQL2 := `SELECT user_id, title, description, status, pending_at, created_at FROM todos WHERE id = $1 AND archived_at IS NULL`
	var todo2 models.TodoClient
	err2 := database.Todo.Get(&todo2, SQL2, id)
	if err2 != nil {
		return models.TodoClient{}, err2
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

func ToggleTodoById(id int) (models.TodoClient, error) {
	var status bool
	SQL := `SELECT status FROM todos WHERE id=$1 AND archived_at IS NULL`
	err := database.Todo.Get(&status, SQL, id)
	if err != nil {
		return models.TodoClient{}, err
	}
	fmt.Println("here 1")
	SQL1 := `UPDATE todos SET status=$1 WHERE id=$2`
	if status == true {
		_, err1 := database.Todo.Exec(SQL1, false, id)
		fmt.Println("here 2")
		if err1 != nil {
			return models.TodoClient{}, err1
		}
	} else {
		_, err1 := database.Todo.Exec(SQL1, true, id)
		fmt.Println("here 3")
		if err1 != nil {
			return models.TodoClient{}, err1
		}
		fmt.Println("here 4")
	}

	SQL2 := `SELECT user_id, title, description, status, pending_at, created_at FROM todos WHERE id = $1 AND archived_at IS NULL`
	var todo models.TodoClient
	err2 := database.Todo.Get(&todo, SQL2, id)
	fmt.Println("here 5")
	if err2 != nil {
		return models.TodoClient{}, err2
	}
	fmt.Println("here 6")
	return todo, nil
}

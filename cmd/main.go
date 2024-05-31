package main

import (
	"04_todo_swagger/database"
	"04_todo_swagger/handlers"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/user/register", handlers.RegisterUser).Methods("POST")
	r.HandleFunc("/user/login/{userId}", handlers.LoginUser).Methods("POST")
	r.HandleFunc("/user/logout/{userId}", handlers.VerifyUserMidd(handlers.LogoutUser)).Methods("POST")

	r.HandleFunc("/user/{userId}/todo", handlers.VerifyUserMidd(handlers.CreateTodo)).Methods("POST")
	r.HandleFunc("/user/{userId}/todo", handlers.VerifyUserMidd(handlers.ReadAllTodo)).Methods("GET")
	r.HandleFunc("/user/{userId}/todo/{todoId}", handlers.VerifyUserMidd(handlers.ReadTodo)).Methods("GET")
	r.HandleFunc("/user/{userId}/todo/{todoId}", handlers.VerifyUserMidd(handlers.UpdateTodo)).Methods("PUT")
	r.HandleFunc("/user/{userId}/todo/{todoId}", handlers.VerifyUserMidd(handlers.DeleteTodo)).Methods("DELETE")

	r.HandleFunc("/user/{userId}/todo/{todoId}/toggle", handlers.VerifyUserMidd(handlers.ToggleTodo)).Methods("PUT")

	if err := database.ConnectAndMigrate(
		"localhost",
		"5440",
		"todo-go",
		"local",
		"local",
		database.SSLModeDisable); err != nil {
		log.Fatal("failed to initialize and migrate database with error: %+v", err)
	}
	fmt.Println("migration successful!!")

	log.Fatal(http.ListenAndServe(":8080", r))
}

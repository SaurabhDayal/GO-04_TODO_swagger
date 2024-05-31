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

	r.HandleFunc("/user/{userId}/todo", handlers.CreateTodo).Methods("POST")
	r.HandleFunc("/user/{userId}/todo", handlers.ReadAllTodo).Methods("GET")
	r.HandleFunc("/user/{userId}/todo/{todoId}", handlers.ReadTodo).Methods("GET")
	r.HandleFunc("/user/{userId}/todo/{todoId}", handlers.UpdateTodo).Methods("PUT")
	r.HandleFunc("/user/{userId}/todo/{todoId}", handlers.DeleteTodo).Methods("DELETE")

	if err := database.ConnectAndMigrate(
		"localhost",
		"5440",
		"todo-go",
		"local",
		"local",
		database.SSLModeDisable); err != nil {
		log.Fatalf("failed to initialize and migrate database with error: %+v", err)
	}
	fmt.Println("migration successful!!")

	log.Fatal(http.ListenAndServe(":8080", r))
}

package main

import (
	"04_todo_swagger/database"
	"04_todo_swagger/handlers"
	"fmt"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
)

// @title Your Project API
// @version 1.0
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
func main() {

	r := mux.NewRouter()

	r.HandleFunc("/user/{userId}/todo", handlers.CreateTodo).Methods("POST")
	r.HandleFunc("/user/{userId}/todo", handlers.ReadAllTask).Methods("GET")
	r.HandleFunc("/user/{userId}/todo/{todoId}", handlers.ReadTask).Methods("GET")
	r.HandleFunc("/user/{userId}/todo/{todoId}", handlers.UpdateTask).Methods("PUT")
	r.HandleFunc("/user/{userId}/todo/{todoId}", handlers.DeleteTask).Methods("DELETE")

	// Serve Swagger UI
	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), // The URL to API definition
	))

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

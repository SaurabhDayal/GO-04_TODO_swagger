package main

import (
	"04_todo_swagger/database"
	"04_todo_swagger/handlers"
	"fmt"
	"github.com/go-chi/chi"
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
	// Initialize the Chi router
	r := chi.NewRouter()

	// Define routes for CRUD operations
	r.Post("/task", handlers.CreateTask)
	r.Get("/task", handlers.ReadAllTask)
	r.Get("/task/{taskId}", handlers.ReadTask)
	r.Put("/task/{taskId}", handlers.UpdateTask)
	r.Delete("/task/{taskId}", handlers.DeleteTask)

	// Serve Swagger UI
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), // URL to API definition
	))

	// Initialize and migrate the database
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

	// Start the HTTP server
	log.Fatal(http.ListenAndServe(":8080", r))
}

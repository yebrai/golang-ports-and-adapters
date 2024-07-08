package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"golang-ports-and-adapters/internal/core/services"
	"golang-ports-and-adapters/internal/handlers"
	"golang-ports-and-adapters/internal/repository"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	todoRepository := repository.NewMySQLTodoRepository(db)
	todoService := services.NewTodoService(todoRepository)
	todoHandler := handlers.NewTodoHandler(todoService)

	r := mux.NewRouter()
	r.HandleFunc("/todo/create", todoHandler.CreateTodoHandler).Methods("POST")
	r.HandleFunc("/todo/edit", todoHandler.UpdateTodoHandler).Methods("PUT")
	r.HandleFunc("/todo/{id:[0-9]+}", todoHandler.GetTodoHandler).Methods("GET")

	//templates
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"golang-ports-and-adapters/internal/core/services"
	"golang-ports-and-adapters/internal/handlers"
	"golang-ports-and-adapters/internal/repository"
)

func main() {
	// Conectar a la base de datos
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/dbname")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Repositorio y servicio
	repo := repository.NewMySQLRepository(db)
	service := services.NewTodoService(repo)

	// Handlers
	todoHandler := handlers.NewTodoHandler(service)

	// Router
	router := mux.NewRouter()
	router.HandleFunc("/todo/create", todoHandler.CreateTodoHandler).Methods("POST")
	router.HandleFunc("/todo/edit", todoHandler.UpdateTodoHandler).Methods("PUT")
	router.HandleFunc("/todo/{id}", todoHandler.GetTodoHandler).Methods("GET")

	// Servidor
	log.Println("Servidor iniciado en :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

package repository

import (
	"database/sql"
	"golang-ports-and-adapters/internal/core/domain"
	"golang-ports-and-adapters/internal/core/ports"
)

type MySQLRepository struct {
	DB *sql.DB
}

func NewMySQLRepository(db *sql.DB) ports.TodoRepository {
	return &MySQLRepository{DB: db}
}

func (r *MySQLRepository) Create(todo *domain.Todo) error {
	_, err := r.DB.Exec("INSERT INTO todos (title, status) VALUES (?, ?)", todo.Title, todo.Status)
	return err
}

func (r *MySQLRepository) Update(todo *domain.Todo) error {
	_, err := r.DB.Exec("UPDATE todos SET title=?, status=? WHERE id=?", todo.Title, todo.Status, todo.ID)
	return err
}

func (r *MySQLRepository) GetByID(id int) (*domain.Todo, error) {
	row := r.DB.QueryRow("SELECT id, title, status FROM todos WHERE id=?", id)
	todo := &domain.Todo{}
	err := row.Scan(&todo.ID, &todo.Title, &todo.Status)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

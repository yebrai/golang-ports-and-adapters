package repository

import (
	"database/sql"
	"golang-ports-and-adapters/internal/core/domain"
)

type MySQLTodoRepository struct {
	db *sql.DB
}

func NewMySQLTodoRepository(db *sql.DB) *MySQLTodoRepository {
	return &MySQLTodoRepository{db: db}
}

func (r *MySQLTodoRepository) Create(todo *domain.Todo) error {
	_, err := r.db.Exec("INSERT INTO todos (title, status) VALUES (?, ?)", todo.Title, todo.Status)
	return err
}

func (r *MySQLTodoRepository) Update(todo *domain.Todo) error {
	_, err := r.db.Exec("UPDATE todos SET title=?, status=? WHERE id=?", todo.Title, todo.Status, todo.ID)
	return err
}

func (r *MySQLTodoRepository) GetByID(id int) (*domain.Todo, error) {
	row := r.db.QueryRow("SELECT id, title, status FROM todos WHERE id=?", id)
	var todo domain.Todo
	err := row.Scan(&todo.ID, &todo.Title, &todo.Status)
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

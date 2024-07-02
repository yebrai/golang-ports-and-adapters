package ports

import "golang-ports-and-adapters/internal/core/domain"

type TodoRepository interface {
	Create(todo *domain.Todo) error
	Update(todo *domain.Todo) error
	GetByID(id int) (*domain.Todo, error)
}

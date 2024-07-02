package services

import (
	"golang-ports-and-adapters/internal/core/domain"
	"golang-ports-and-adapters/internal/core/ports"
)

type TodoService struct {
	repo ports.TodoRepository
}

func NewTodoService(repo ports.TodoRepository) *TodoService {
	return &TodoService{repo: repo}
}

func (s *TodoService) Create(todo *domain.Todo) error {
	return s.repo.Create(todo)
}

func (s *TodoService) Update(todo *domain.Todo) error {
	return s.repo.Update(todo)
}

func (s *TodoService) GetByID(id int) (*domain.Todo, error) {
	return s.repo.GetByID(id)
}

package service

import (
	"fmt"
	"go_crud/internal/repository"
)

type TodoService struct {
	repo repository.TodoRepository
}

func NewTodoService(repo repository.TodoRepository) *TodoService {
	return &TodoService{repo: repo}
}

func (s *TodoService) ListTodos() ([]repository.Todo, error) {
	return s.repo.List()
}

func (s *TodoService) CreateTodo(title string) (repository.Todo, error) {
	t, err := s.repo.Create(title)
	if err != nil {
		return repository.Todo{}, fmt.Errorf("create todo: %w", err)
	}
	return t, nil
}

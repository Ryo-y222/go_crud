package service

import "go_crud/internal/repository"

type TodoService struct {
	repo repository.TodoRepository
}

func NewTodoService(repo repository.TodoRepository) *TodoService {
	return &TodoService{repo: repo}
}

func (s *TodoService) ListTodos() ([]repository.Todo, error) {
	return s.repo.List()
}

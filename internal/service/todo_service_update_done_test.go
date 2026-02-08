package service

import (
	"database/sql"
	"errors"
	"testing"

	"go_crud/internal/repository"
)

type fakeTodoRepo2 struct {
	updateDoneFn func(id int64, done bool) error
}

// interface合わせのための最小stub（必要に応じて型だけ直す）
func (f *fakeTodoRepo2) List() ([]repository.Todo, error) { return nil, nil }
func (f *fakeTodoRepo2) Create(title string) (repository.Todo, error) {
	return repository.Todo{}, nil
}
func (f *fakeTodoRepo2) Delete(id int64) error { return nil }

func (f *fakeTodoRepo2) UpdateDone(id int64, done bool) error {
	if f.updateDoneFn == nil {
		return nil
	}
	return f.updateDoneFn(id, done)
}

func TestTodoService_UpdateTodoDone(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		s := NewTodoService(&fakeTodoRepo2{
			updateDoneFn: func(id int64, done bool) error { return nil },
		})
		if err := s.UpdateTodoDone(1, true); err != nil {
			t.Fatalf("expected nil, got %v", err)
		}
	})

	t.Run("not found -> ErrTodoNotFound", func(t *testing.T) {
		s := NewTodoService(&fakeTodoRepo2{
			updateDoneFn: func(id int64, done bool) error { return sql.ErrNoRows },
		})
		err := s.UpdateTodoDone(999, true)
		if !errors.Is(err, repository.ErrTodoNotFound) {
			t.Fatalf("expected ErrTodoNotFound, got %v", err)
		}
	})

	t.Run("repo error is wrapped", func(t *testing.T) {
		repoErr := errors.New("db down")
		s := NewTodoService(&fakeTodoRepo2{
			updateDoneFn: func(id int64, done bool) error { return repoErr },
		})
		err := s.UpdateTodoDone(1, true)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !errors.Is(err, repoErr) {
			t.Fatalf("expected wrapped repoErr, got %v", err)
		}
	})
}

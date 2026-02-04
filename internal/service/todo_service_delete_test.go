package service

import (
	"database/sql"
	"errors"
	"testing"

	"go_crud/internal/repository"
)

// fake repo: Deleteだけ本物、他はコンパイルを通すための最小stub
type fakeTodoRepo struct {
	deleteFn func(id int64) error
}

// List / Create の戻り値型が違う場合は、todo_repository.go の定義に合わせて修正
func (f *fakeTodoRepo) List() ([]repository.Todo, error) { // 例：repository.Todo が無いなら型を合わせる
	return nil, nil
}
func (f *fakeTodoRepo) Create(title string) (repository.Todo, error) { // 例
	return repository.Todo{}, nil
}
func (f *fakeTodoRepo) UpdateDone(id int64, done bool) error {
	return nil
}

// ↑↑↑ ここまで interface合わせゾーン ↑↑↑

func (f *fakeTodoRepo) Delete(id int64) error {
	if f.deleteFn == nil {
		return nil
	}
	return f.deleteFn(id)
}

func TestTodoService_Delete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		s := NewTodoService(&fakeTodoRepo{
			deleteFn: func(id int64) error { return nil },
		})

		if err := s.Delete(1); err != nil {
			t.Fatalf("expected nil, got %v", err)
		}
	})

	t.Run("not found -> ErrTodoNotFound", func(t *testing.T) {
		s := NewTodoService(&fakeTodoRepo{
			deleteFn: func(id int64) error { return sql.ErrNoRows },
		})

		err := s.Delete(999)
		if !errors.Is(err, repository.ErrTodoNotFound) {
			t.Fatalf("expected ErrTodoNotFound, got %v", err)
		}
	})

	t.Run("repo error is wrapped", func(t *testing.T) {
		repoErr := errors.New("db down")
		s := NewTodoService(&fakeTodoRepo{
			deleteFn: func(id int64) error { return repoErr },
		})

		err := s.Delete(1)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !errors.Is(err, repoErr) {
			t.Fatalf("expected wrapped repoErr, got %v", err)
		}
	})
}

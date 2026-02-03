package repository

type TodoRepository interface {
	List() ([]Todo, error)
}

// 開発用（DB繋ぐまでの仮実装）
type InMemoryTodoRepository struct{}

func NewInMemoryTodoRepository() *InMemoryTodoRepository {
	return &InMemoryTodoRepository{}
}

func (r *InMemoryTodoRepository) List() ([]Todo, error) {
	return []Todo{}, nil
}

package repository

type TodoRepository interface {
	List() ([]Todo, error)
	Create(title string) (Todo, error)
	UpdateDone(id int64, done bool) error
}

// 開発用（DB繋ぐまでの仮実装）
type InMemoryTodoRepository struct{}

func NewInMemoryTodoRepository() *InMemoryTodoRepository {
	return &InMemoryTodoRepository{}
}

func (r *InMemoryTodoRepository) List() ([]Todo, error) {
	return []Todo{}, nil
}

func (r *InMemoryTodoRepository) Create(title string) (Todo, error) {
	return Todo{ID: 1, Title: title, Done: false}, nil
}

func (r *InMemoryTodoRepository) UpdateDone(id int64, done bool) error {
	return nil
}

package repository

type Todo struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type TodoRepository interface {
	List() ([]Todo, error)
}

type InMemoryTodoRepository struct{}

func NewInMemoryTodoRepository() *InMemoryTodoRepository {
	return &InMemoryTodoRepository{}
}

func (r *InMemoryTodoRepository) List() ([]Todo, error) {
	return []Todo{}, nil
}

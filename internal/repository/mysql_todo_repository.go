package repository

import "database/sql"

type MySQLTodoRepository struct {
	db *sql.DB
}

func NewMySQLTodoRepository(db *sql.DB) *MySQLTodoRepository {
	return &MySQLTodoRepository{db: db}
}

func (r *MySQLTodoRepository) List() ([]Todo, error) {
	rows, err := r.db.Query(`SELECT id, title, done FROM todos ORDER BY id DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := make([]Todo, 0)
	for rows.Next() {
		var t Todo
		if err := rows.Scan(&t.ID, &t.Title, &t.Done); err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *MySQLTodoRepository) Create(title string) (Todo, error) {
	res, err := r.db.Exec(`INSERT INTO todos (title, done) VALUES (?, false)`, title)
	if err != nil {
		return Todo{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return Todo{}, err
	}

	// 最低限：作成直後の状態を返す
	return Todo{
		ID:    id,
		Title: title,
		Done:  false,
	}, nil
}

func (r *MySQLTodoRepository) UpdateDone(id int64, done bool) error {

	res, err := r.db.Exec(`UPDATE todos SET done = ? WHERE id = ?`, done, id)

	if err != nil {
		return err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if n == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (r *MySQLTodoRepository) Delete(id int64) error {
	result, err := r.db.Exec(`DELETE FROM todos WHERE id=?`, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return sql.ErrNoRows
	}
	return nil
}

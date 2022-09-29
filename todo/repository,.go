package todo

import (
	"database/sql"
)

type DB interface {
	QueryRow(query string, args ...any) *sql.Row
	Query(query string, args ...any) (*sql.Rows, error)
}

type Repository struct {
	db DB
}

func NewRepository(db DB) *Repository {
	return &Repository{db: db}
}
func (r *Repository) AddTodo(todo Todo) error {
	query := `
	INSERT INTO todos (id, content)
	VALUES ($1, $2)`
	row := r.db.QueryRow(query, todo.ID, todo.Content)
	if err := row.Err(); err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetTodos() ([]Todo, error) {
	query := `
	SELECT id, content from todos ORDER BY id DESC`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	todos := make([]Todo, 0)
	for rows.Next() {
		var todo Todo
		err = rows.Scan(
			&todo.ID, &todo.Content)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	if err = rows.Close(); err != nil {
		return nil, err
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return todos, nil
}
func (r *Repository) HasTodo(content string) (bool, error) {
	query := `
	SELECT EXISTS (SELECT 1 from todos t WHERE t.content = $1)`

	row := r.db.QueryRow(query, content)
	var hasTodo bool
	if err := row.Scan(&hasTodo); err != nil {
		return false, err
	}
	return hasTodo, nil
}

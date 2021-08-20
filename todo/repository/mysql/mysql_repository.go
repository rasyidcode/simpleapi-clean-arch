package mysql

import (
	"context"
	"database/sql"

	"rasyidcode/simpleapi-clean-arch/models"
)

type mysqlTodoRepo struct {
	DB *sql.DB
}

func NewMysqlTodoRepository(db *sql.DB) models.TodoRepository {
	return &mysqlTodoRepo{
		DB: db,
	}
}

func (m *mysqlTodoRepo) GetOne(ctx context.Context, query string, args ...interface{}) (res models.Todo, err error) {
	stmt, err := m.DB.PrepareContext(ctx, query)
	if err != nil {
		return models.Todo{}, err
	}

	row := stmt.QueryRowContext(ctx, args...)
	res = models.Todo{}

	err = row.Scan(
		&res.ID,
		&res.TodoName,
		&res.Desc,
	)
	return
}

func (m *mysqlTodoRepo) GetByID(ctx context.Context, id int64) (models.Todo, error) {
	query := `SELECT id, todo_name, desc FROM todo WHERE id=?`
	return m.GetOne(ctx, query, id)
}

package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"
)

type Repository interface {
	CreateTodo(ctx context.Context, todo TodoEntity) error
	UpdateTodo(ctx context.Context, todo TodoEntity) error
	DeleteTodo(ctx context.Context, id uuid.UUID) error
	GetTodo(ctx context.Context, id uuid.UUID) (TodoEntity, error)
}

type TodoRepository struct {
	DB     *sql.DB
	Logger echo.Logger
}

func (t *TodoRepository) CreateTodo(ctx context.Context, todo TodoEntity) error {
	result, err := t.DB.ExecContext(ctx, createTodoQuery, todo.ID, todo.Text, todo.CreatedAt, todo.UpdatedAt, todo.Deleted)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		t.Logger.Errorf("no rows were inserted when create todo was called")
		return fmt.Errorf("no todos created")
	}

	return nil
}

func (t *TodoRepository) UpdateTodo(ctx context.Context, todo TodoEntity) error {
	result, err := t.DB.ExecContext(ctx, updateTodoQuery, todo.Text, todo.UpdatedAt, todo.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		t.Logger.Errorf("no rows were updated when update todo was called")
		return fmt.Errorf("no todos updated")
	}

	return nil
}

func (t *TodoRepository) DeleteTodo(ctx context.Context, id uuid.UUID) error {
	result, err := t.DB.ExecContext(ctx, deleteTodoQuery, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		t.Logger.Errorf("no rows were updated when delete todo was called")
		return fmt.Errorf("no todos deleted")
	}

	return nil
}

func (t *TodoRepository) GetTodo(ctx context.Context, id uuid.UUID) (todo TodoEntity, err error) {
	row := t.DB.QueryRowContext(ctx, getTodoQuery, id)
	if row.Err() != nil {
		return todo, row.Err()
	}

	err = row.Scan(&todo.ID, &todo.Text)
	if err != nil {
		return todo, err
	}

	return todo, nil
}

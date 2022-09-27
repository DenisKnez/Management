package repository

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/mock"
)

type MockTodoRepository struct {
	mock.Mock
}

func (t *MockTodoRepository) CreateTodo(ctx context.Context, todo TodoEntity) error {
	args := t.Called(todo)
	return args.Error(0)
}

func (t *MockTodoRepository) UpdateTodo(ctx context.Context, todo TodoEntity) error {
	args := t.Called(todo)
	return args.Error(0)
}

func (t *MockTodoRepository) DeleteTodo(ctx context.Context, id uuid.UUID) error {
	args := t.Called(id)
	return args.Error(0)
}

func (t *MockTodoRepository) GetTodo(ctx context.Context, id uuid.UUID) (todo TodoEntity, err error) {
	args := t.Called(id)
	return args.Get(0).(TodoEntity), args.Error(1)
}

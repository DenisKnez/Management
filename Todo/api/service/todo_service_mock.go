package service

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/mock"
)

type MockTodoService struct {
	mock.Mock
}

func (t *MockTodoService) CreateTodo(ctx context.Context, text string) error {
	args := t.Called(text)
	return args.Error(0)
}

func (t *MockTodoService) UpdateTodo(ctx context.Context, todo Todo) error {
	args := t.Called(todo)
	return args.Error(0)
}

func (t *MockTodoService) DeleteTodo(ctx context.Context, id uuid.UUID) error {
	args := t.Called(id)
	return args.Error(0)
}

func (t *MockTodoService) GetTodo(ctx context.Context, id uuid.UUID) (Todo, error) {
	args := t.Called(id)
	return args.Get(0).(Todo), args.Error(1)
}

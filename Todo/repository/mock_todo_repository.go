package repository

import (
	"context"
	"time"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/mock"
)

type MockTodoRepository struct {
	mock.Mock
}

func (t *MockTodoRepository) CreateTodo(ctx context.Context, todo TodoEntity) error {
	args := t.Called()
	return args.Error(0)
}

func (t *MockTodoRepository) UpdateTodo(ctx context.Context, todo TodoEntity) error {
	return nil
}

func (t *MockTodoRepository) DeleteTodo(ctx context.Context, id uuid.UUID) error {
	return nil
}

func (t *MockTodoRepository) GetTodo(ctx context.Context, id uuid.UUID) (todo TodoEntity, err error) {
	args := t.Called()

	if id == uuid.Must(uuid.FromString("31b9d845-dc19-4f12-bf97-d8b101700be0")) {
		return TodoEntity{
			ID:        id,
			Text:      "test text",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Deleted:   false,
		}, nil

	} else {
		return args.Get(0).(TodoEntity), args.Error(1)
	}
}

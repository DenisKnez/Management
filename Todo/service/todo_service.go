package service

import (
	"context"
	"time"

	"github.com/DenisKnez/management/todo/repository"
	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"
)

type Service interface {
	CreateTodo(ctx context.Context, todo Todo) error
	UpdateTodo(ctx context.Context, id uuid.UUID, todo Todo) error
	DeleteTodo(ctx context.Context, id uuid.UUID) error
	GetTodo(ctx context.Context, id uuid.UUID) (Todo, error)
}

type TodoService struct {
	TodoRepo repository.Repository
	Logger   echo.Logger
}

func (service *TodoService) CreateTodo(ctx context.Context, todo Todo) error {
	now := time.Now()
	id, err := uuid.NewV4()
	if err != nil {
		service.Logger.Errorf("failed to create id for new todo: %v", err)
		return err
	}

	err = service.TodoRepo.CreateTodo(ctx, repository.TodoEntity{
		ID:        id,
		Text:      todo.Text,
		CreatedAt: now,
		UpdatedAt: now,
		Deleted:   false,
	})
	if err != nil {
		service.Logger.Errorf("failed to create todo in database: %v", err)
		return err
	}

	return nil
}

func (t *TodoService) UpdateTodo(ctx context.Context, id uuid.UUID, todo Todo) error {
	now := time.Now()

	err := t.TodoRepo.UpdateTodo(ctx, repository.TodoEntity{
		ID:        id,
		Text:      todo.Text,
		UpdatedAt: now,
	})
	if err != nil {
		t.Logger.Errorf("failed to update todo %v: %v", id, err)
		return err
	}

	return nil
}

func (t *TodoService) DeleteTodo(ctx context.Context, id uuid.UUID) error {
	err := t.TodoRepo.DeleteTodo(ctx, id)
	if err != nil {
		t.Logger.Errorf("failed to delete todo %v: %v", id, err)
		return err
	}

	return nil
}

func (t *TodoService) GetTodo(ctx context.Context, id uuid.UUID) (todo Todo, err error) {
	todoEntity, err := t.TodoRepo.GetTodo(ctx, id)
	if err != nil {
		t.Logger.Errorf("failed to get todo %v: %v", id, err)
		return todo, err
	}

	return Todo{
		ID:   todoEntity.ID,
		Text: todoEntity.Text,
	}, nil
}

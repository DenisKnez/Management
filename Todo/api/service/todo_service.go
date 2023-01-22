package service

import (
	"context"

	"github.com/DenisKnez/management/todo/api/repository"
	"github.com/DenisKnez/util"
	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"
)

type Service interface {
	CreateTodo(ctx context.Context, text string) error
	UpdateTodo(ctx context.Context, todo Todo) error
	DeleteTodo(ctx context.Context, id uuid.UUID) error
	GetTodo(ctx context.Context, id uuid.UUID) (Todo, error)
}

type TodoService struct {
	Util     util.Util
	TodoRepo repository.Repository
	Logger   echo.Logger
}

func (service *TodoService) CreateTodo(ctx context.Context, text string) error {
	id, err := service.Util.UUID().GetNewUUID()
	if err != nil {
		return err
	}

	err = service.TodoRepo.CreateTodo(ctx, repository.TodoEntity{
		ID:        id,
		Text:      text,
		CreatedAt: service.Util.Time().GetCurrentTime(),
		UpdatedAt: service.Util.Time().GetCurrentTime(),
		Deleted:   false,
	})
	if err != nil {
		service.Logger.Errorf("failed to create todo in database: %v", err)
		return err
	}

	return nil
}

func (t *TodoService) UpdateTodo(ctx context.Context, todo Todo) error {
	err := t.TodoRepo.UpdateTodo(ctx, repository.TodoEntity{
		ID:        todo.ID,
		Text:      todo.Text,
		UpdatedAt: t.Util.Time().GetCurrentTime(),
	})
	if err != nil {
		t.Logger.Errorf("failed to update todo %v: %v", todo.ID, err)
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

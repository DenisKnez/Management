package service

import (
	"context"
	"testing"

	"github.com/DenisKnez/management/todo/repository"
	"github.com/gofrs/uuid"
)

func TestCreateTodo(t *testing.T) {
	todoRepo := new(repository.MockTodoRepository)

	todoRepo.On("CreateTodo").Return(nil)

	service := TodoService{
		TodoRepo: todoRepo,
	}

	service.CreateTodo(context.Background(), Todo{
		Text: "hello",
	})

	todoRepo.AssertExpectations(t)
}

func TestGetTodo(t *testing.T) {
	todoRepo := new(repository.MockTodoRepository)
	todo := repository.TodoEntity{Text: "test"}

	todoRepo.On("GetTodo").Return(todo, nil)

	service := TodoService{
		TodoRepo: todoRepo,
	}

	service.GetTodo(context.Background(), uuid.Must(uuid.FromString("31b9d845-dc19-4f12-bf97-d8b101700be1")))

	todoRepo.AssertExpectations(t)
}

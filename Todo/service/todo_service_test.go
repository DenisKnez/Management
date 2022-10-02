package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/DenisKnez/management/todo/repository"
	"github.com/DenisKnez/util"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/require"
)

func TestCreateTodo(t *testing.T) {
	now, err := time.Parse("2006-01-02T15:04:05-0700", "2022-01-02T15:04:05-0700")
	if err != nil {
		t.FailNow()
	}

	todoRepo := new(repository.MockTodoRepository)
	id := uuid.Must(uuid.FromString("037d562a-f5f9-41df-9ded-f44991dd2309"))
	util := util.New(&now, &id)

	todoRepo.On("CreateTodo", repository.TodoEntity{
		ID:        id,
		Text:      "test",
		CreatedAt: now,
		UpdatedAt: now,
		Deleted:   false,
	}).Return(nil)

	service := TodoService{
		Util:     *util,
		TodoRepo: todoRepo,
	}

	service.CreateTodo(context.Background(), Todo{
		Text: "test",
	})

	todoRepo.AssertExpectations(t)
}

func TestUpdateTodo(t *testing.T) {
	now, err := time.Parse("2006-01-02T15:04:05-0700", "2022-01-02T15:04:05-0700")
	if err != nil {
		t.FailNow()
	}
	id := uuid.Must(uuid.FromString("037d562a-f5f9-41df-9ded-f44991dd2309"))
	todoEntity := repository.TodoEntity{ID: id, Text: "test", UpdatedAt: now}

	todoRepo := new(repository.MockTodoRepository)
	todoRepo.On("UpdateTodo", todoEntity).Return(nil)

	util := util.New(&now, &id)

	service := TodoService{
		Util:     *util,
		TodoRepo: todoRepo,
	}

	err = service.UpdateTodo(context.Background(), Todo{
		ID:   id,
		Text: "test",
	})
	todoRepo.AssertExpectations(t)

	require.Nil(t, err)
}

func TestDeleteTodo(t *testing.T) {
	now, err := time.Parse("2006-01-02T15:04:05-0700", "2022-01-02T15:04:05-0700")
	if err != nil {
		t.FailNow()
	}
	id := uuid.Must(uuid.FromString("037d562a-f5f9-41df-9ded-f44991dd2309"))

	todoRepo := new(repository.MockTodoRepository)
	todoRepo.On("DeleteTodo", id).Return(nil)

	util := util.New(&now, &id)

	service := TodoService{
		Util:     *util,
		TodoRepo: todoRepo,
	}

	err = service.DeleteTodo(context.Background(), id)
	todoRepo.AssertExpectations(t)

	require.Nil(t, err)
}

func TestGetTodo(t *testing.T) {
	id := uuid.Must(uuid.FromString("037d562a-f5f9-41df-9ded-f44991dd2309"))
	badId := uuid.Must(uuid.FromString("037d562a-f5f9-41df-9ded-f44991dd2309"))
	todoRepo := new(repository.MockTodoRepository)
	todoEntity := repository.TodoEntity{ID: id, Text: "test"}

	mockCall := todoRepo.On("GetTodo", id).Return(todoEntity, nil)

	service := TodoService{
		TodoRepo: todoRepo,
	}

	todo, err := service.GetTodo(context.Background(), id)
	todoRepo.AssertExpectations(t)

	require.Nil(t, err)
	require.Equal(t, todo.ID, id)

	mockCall.Unset()

	todoRepo.On("GetTodo", badId).Return(todo, errors.New("failed"))

	todoRepo.AssertExpectations(t)
}

package grpc

import (
	"context"
	"fmt"

	"github.com/DenisKnez/management/todo/service"
	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"
)

type GrpcHandler struct {
	UnimplementedTodoServer
	TodoService service.Service
	Logger      echo.Logger
}

func (gh *GrpcHandler) CreateTodo(ctx context.Context, req *CreateTodoRequest) (*CreateTodoResponse, error) {
	err := gh.TodoService.CreateTodo(ctx, service.Todo{
		Text: req.Text,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create todo")
	}

	return &CreateTodoResponse{}, nil
}

func (gh *GrpcHandler) DeleteTodo(ctx context.Context, req *DeleteTodoRequest) (*DeleteTodoResponse, error) {
	id, err := uuid.FromString(req.Id)
	if err != nil {
		gh.Logger.Errorf("failed to convert request id %s to uuid: %v", req.Id, err)
		return nil, fmt.Errorf("provided is an invalid uuid format")
	}

	err = gh.TodoService.DeleteTodo(ctx, id)
	if err != nil {
		gh.Logger.Errorf("failed to delete todo: %v", err)
		return nil, fmt.Errorf("failed to delete todo")
	}

	return &DeleteTodoResponse{}, nil
}

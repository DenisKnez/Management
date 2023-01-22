package service

import (
	"context"
	"testing"

	"github.com/DenisKnez/management/user/repository"
	"github.com/DenisKnez/management/user/service/grpc"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	repo := &repository.MockUserRepository{}

	todoServiceClient := &grpc.MockTodoServiceClient{}
	todoServiceClient.On("CreateTodo", &grpc.CreateTodoRequest{
		Text: "create user todo",
	}).Return(&grpc.CreateTodoResponse{}, nil)

	service := UserService{
		UserRepo:          repo,
		TodoServiceClient: todoServiceClient,
	}

	repo.On("CreateUser", repository.UserEntity{
		Name: "test",
	}).Return(nil)

	service.CreateUser(context.Background(), User{
		Name: "test",
	})
}

func TestUploadFile(t *testing.T) {
	repo := &repository.MockUserRepository{}
	todoService := &grpc.MockTodoServiceClient{}
	service := UserService{
		UserRepo:          repo,
		TodoServiceClient: todoService,
	}

	todoService.On("UploadFile").Return()

	err := service.UploadFile(context.Background(), File{
		Name: "test",
		Data: []byte("hello"),
	})

	repo.AssertExpectations(t)
	require.Nil(t, err)
}

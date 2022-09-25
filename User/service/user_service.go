package service

import (
	"context"

	"github.com/DenisKnez/management/user/repository"
	userGrpc "github.com/DenisKnez/management/user/service/grpc"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
	CreateUser(ctx context.Context, user User) error
	UpdateUser(ctx context.Context, id string, user User) error
	// DeleteUser(ctx context.Context, id uuid.UUID) error
	// GetUser(ctx context.Context, id uuid.UUID) (User, error)
	// GetUsers(ctx context.Context) ([]User, error)
}

type UserService struct {
	UserRepo          repository.Repository
	TodoServiceClient userGrpc.TodoClient
}

func (service *UserService) CreateUser(ctx context.Context, user User) error {
	_, err := service.TodoServiceClient.CreateTodo(ctx, &userGrpc.CreateTodoRequest{
		Text: "create user todo",
	})
	if err != nil {
		return err
	}

	return service.UserRepo.CreateUser(ctx, repository.UserEntity{
		Name: user.Name,
	})
}

func (service *UserService) UpdateUser(ctx context.Context, id string, user User) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	return service.UserRepo.UpdateUser(ctx, repository.UserEntity{
		ID:   objectID,
		Name: user.Name,
	})
}

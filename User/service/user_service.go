package service

import (
	"context"
	"fmt"

	"github.com/DenisKnez/management/user/repository"
	userGrpc "github.com/DenisKnez/management/user/service/grpc"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
	CreateUser(ctx context.Context, user User) error
	UpdateUser(ctx context.Context, id string, user User) error
	UploadFile(ctx context.Context, file File) error
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

func (service *UserService) UploadFile(ctx context.Context, file File) error {
	client, err := service.TodoServiceClient.UploadFile(ctx)
	if err != nil {
		return err
	}

	if len(file.Data) <= 10_000_000 {
		sendFile(client, file.Name, file.Data)
		return nil
	}

	sections := len(file.Data) / 10_000_000

	for i := 0; i <= sections; i++ {
		var data []byte
		if i == sections {
			data = file.Data[(i-1)*10_000_000:]
		} else {
			data = file.Data[i : i*10_000_000]
		}

		sendFile(client, file.Name, data)

		fmt.Println("transfering 10 megabytes...")
	}
	return nil
}

func sendFile(client userGrpc.Todo_UploadFileClient, fileName string, fileData []byte) error {
	err := client.Send(&userGrpc.UploadFileRequest{
		FileName: fileName,
		FileData: fileData,
	})
	if err != nil {
		return err
	}

	return nil
}

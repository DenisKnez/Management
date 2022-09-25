package main

import (
	"github.com/DenisKnez/management/user/handler"
	"github.com/DenisKnez/management/user/repository"
	"github.com/DenisKnez/management/user/service"
	userGrpc "github.com/DenisKnez/management/user/service/grpc"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func setupUser(g *gin.Engine, mongoDB *mongo.Client, todoServiceClient *userGrpc.TodoClient) {

	repo := &repository.UserRepository{
		MongoDB: mongoDB.Database("user"),
	}

	userService := service.UserService{
		UserRepo:          repo,
		TodoServiceClient: *todoServiceClient,
	}

	userHandler := handler.UserHandler{
		UserService: &userService,
	}

	g.POST("/users", userHandler.CreateUser)
	g.POST("/users/:id", userHandler.UpdateUser)
}

package main

import (
	"database/sql"

	todoGrpc "github.com/DenisKnez/management/todo/api/handler/grpc"
	"github.com/DenisKnez/management/todo/api/handler/http"
	"github.com/DenisKnez/management/todo/api/repository"
	"github.com/DenisKnez/management/todo/api/service"
	"github.com/DenisKnez/util"
	"github.com/labstack/echo/v4"
)

func setupHandlers(db *sql.DB, logger echo.Logger) (*http.TodoHandler, *todoGrpc.GrpcHandler) {
	repo := repository.TodoRepository{
		DB:     db,
		Logger: logger,
	}

	todoService := service.TodoService{
		TodoRepo: &repo,
		Logger:   logger,
		Util:     *util.New(nil, nil),
	}

	todoHttpHandler := http.TodoHandler{
		TodoService: &todoService,
		Logger:      logger,
	}

	todoGrpcHandler := todoGrpc.GrpcHandler{
		TodoService: &todoService,
		Logger:      logger,
	}

	return &todoHttpHandler, &todoGrpcHandler
}

func (server *HttpServer) setupRoutes(todoHttpHandler *http.TodoHandler) {
	group := server.Echo.Group("/todos")

	_ = group.POST("", todoHttpHandler.CreateTodo)
	_ = group.POST("/:id", todoHttpHandler.UpdateTodo)
	_ = group.DELETE("/:id", todoHttpHandler.DeleteTodo)
	_ = group.GET("/:id", todoHttpHandler.GetTodo)
}

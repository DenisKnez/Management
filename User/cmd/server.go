package main

import (
	"context"
	"fmt"
	"time"

	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	userGrpc "github.com/DenisKnez/management/user/service/grpc"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Server struct {
}

func (server *Server) dialTodoService(uri string) (*grpc.ClientConn, error) {
	return grpc.Dial("todo:50001", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
}

func (server *Server) createTodoClient(conn *grpc.ClientConn) *userGrpc.TodoClient {
	todoServiceClient := userGrpc.NewTodoClient(conn)
	return &todoServiceClient
}

func (server *Server) connectToMongo(uri string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// TODO: get the number of attempts from config
	attempts := 10

	for i := 0; i < attempts; i++ {
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
		if err != nil {
			log.Printf("attempt to connect to mongo db failed: %v", err)
			continue
		}

		err = client.Ping(ctx, readpref.Primary())
		if err != nil {
			log.Printf("failed to ping database")
			continue
		}

		time.Sleep(2 * time.Second)
		log.Printf("connected to mongo db!")
		return client, nil
	}

	return nil, fmt.Errorf("failed to connect to database")
}

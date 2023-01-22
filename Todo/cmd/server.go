package main

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	todoGrpc "github.com/DenisKnez/management/todo/api/handler/grpc"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

//go:embed db/migrations/*.sql
var migrations embed.FS

type HttpServer struct {
	Echo   *echo.Echo
	Config *Config
}

func NewServer(echo *echo.Echo, config *Config) *HttpServer {
	return &HttpServer{
		Echo:   echo,
		Config: config,
	}
}

func (server *HttpServer) handleGracefulShutdown() {
	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Echo.Shutdown(ctx); err != nil {
		server.Echo.Logger.Fatal(err)
	}
}

func (server *HttpServer) httpListen() {
	bb := fmt.Sprintf(":%s", server.Config.HttpPort)
	server.Echo.Logger.Infof("Starting http listener on port: %s", bb)
	if err := server.Echo.Start(bb); err != nil && err != http.ErrServerClosed {
		server.Echo.Logger.Fatalf("failed to start http server: %v", err)
	}
}

func (server *HttpServer) grpcListen(grpcHandler *todoGrpc.GrpcHandler) {
	ss := fmt.Sprintf(":%s", server.Config.GrpcPort)
	fmt.Println("the grpc port: ", ss)
	listener, err := net.Listen("tcp", ss)
	if err != nil {
		server.Echo.Logger.Fatalf("failed to start grpc listener, shutting down... [Error: %v]\n", err)
	}

	grpcServer := grpc.NewServer()
	todoGrpc.RegisterTodoServer(grpcServer, grpcHandler)

	server.Echo.Logger.Infof("Starting grpc listener on port: %s", server.Config.GrpcPort)
	err = grpcServer.Serve(listener)
	if err != nil || err != grpc.ErrServerStopped {
		server.Echo.Logger.Fatalf("failed to serve grpc server: %v\n", err)
	}
}

func (server *HttpServer) migratePostgres() {
	migs, err := iofs.New(migrations, "db/migrations")
	if err != nil {
		server.Echo.Logger.Fatalf("failed to initialize migrator: %v", err)
	}

	m, err := migrate.NewWithSourceInstance("iofs", migs, server.Config.PostgresDSN)
	if err != nil {
		server.Echo.Logger.Fatalf("migrator failed to connect to database: %v", err)
	}

	defer m.Close()

	err = m.Up()
	if err != nil {
		if err.Error() == "no change" {
			server.Echo.Logger.Info("no migrations were applied")
		} else {
			server.Echo.Logger.Fatalf("failed to migrate database: %v", err)
		}
	}
}

func (server *HttpServer) connectToPostgres() *sql.DB {
	for i := 0; i < server.Config.NumOfAttemptsToConnectToDatabase; i++ {
		connection, err := openPostgresDB(server.Config.PostgresDSN)
		if err != nil {
			server.Echo.Logger.Errorf("attempt to connect to postgres failed: %v", err)
		} else {
			server.Echo.Logger.Info("Connected to postgres")
			return connection
		}

		time.Sleep(2 * time.Second)
	}

	return nil
}

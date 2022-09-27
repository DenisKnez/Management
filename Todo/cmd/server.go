package main

import (
	"context"
	"database/sql"
	"embed"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	todoGrpc "github.com/DenisKnez/management/todo/handler/grpc"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoLog "github.com/labstack/gommon/log"
	"google.golang.org/grpc"
)

//go:embed db/migrations/*.sql
var migrations embed.FS

type HttpServer struct {
	Echo *echo.Echo
}

func NewServer(echo *echo.Echo) *HttpServer {
	return &HttpServer{
		Echo: echo,
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
	httpPort := ":4141"
	server.Echo.Logger.Infof("Starting http listener on port: %s", httpPort)
	if err := server.Echo.Start(httpPort); err != nil && err != http.ErrServerClosed {
		server.Echo.Logger.Fatalf("failed to start http server: %v", err)
	}
}

func (server *HttpServer) grpcListen(grpcHandler *todoGrpc.GrpcHandler) {
	grpcPort := ":50001"
	listener, err := net.Listen("tcp", grpcPort)
	if err != nil {
		server.Echo.Logger.Fatal("failed to start grpc listener, shutting down...")
	}

	grpcServer := grpc.NewServer()
	todoGrpc.RegisterTodoServer(grpcServer, grpcHandler)

	server.Echo.Logger.Infof("Starting grpc listener on port: %s", grpcPort)
	err = grpcServer.Serve(listener)
	if err != nil || err != grpc.ErrServerStopped {
		server.Echo.Logger.Fatalf("failed to start grpc server: %v\n", err)
	}
}

func (server *HttpServer) migratePostgres() {
	migs, err := iofs.New(migrations, "db/migrations")
	if err != nil {
		server.Echo.Logger.Fatalf("failed to initialize migrator: %v", err)
	}

	m, err := migrate.NewWithSourceInstance("iofs", migs, "postgres://postgres:notebook@localhost:5432/todo?sslmode=disable")
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
	// TODO: get the number of attempts from config
	attempts := 10

	for i := 0; i < attempts; i++ {
		connection, err := openPostgresDB("host=localhost port=5432 user=postgres password=notebook dbname=todo sslmode=disable")
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

func (server *HttpServer) setupEchoMiddleware() {
	server.Echo.Use(middleware.Logger())
	server.Echo.Use(middleware.Recover())
}

func (server *HttpServer) setupEchoLogger() {
	server.Echo.Logger.SetLevel(echoLog.INFO)
}

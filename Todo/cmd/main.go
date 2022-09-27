package main

import (
	_ "github.com/DenisKnez/management/todo/docs"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Todo
// @version 1.0
// @description Todo API

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @host localhost:4141
// @BasePath /
func main() {
	echo := echo.New()
	server := NewServer(echo)

	db := server.connectToPostgres()
	if db == nil {
		server.Echo.Logger.Fatal("failed to connect to postgres, shutting down...")
	}

	server.migratePostgres()

	// Serve swagger docs
	// swagger/index.html
	server.Echo.GET("/swagger/*", echoSwagger.WrapHandler)

	// Echo setup
	server.setupEchoLogger()
	server.setupEchoMiddleware()

	httpHandler, grpcHandler := setupHandlers(db, echo.Logger)
	server.setupRoutes(httpHandler)

	// Start server
	go server.httpListen()
	go server.grpcListen(grpcHandler)
	server.handleGracefulShutdown()
}

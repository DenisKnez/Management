package main

import (
	"fmt"
	"os"

	_ "github.com/DenisKnez/management/todo/docs"

	"github.com/go-ini/ini"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

const configPathEnvName = "CONFIG_PATH"
const swaggerURL = "/swagger/*"

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

	// Echo setup
	setupEchoLogger(echo)
	setupEchoMiddleware(echo)

	configPath := os.Getenv(configPathEnvName)
	if configPath == "" {
		echo.Logger.Fatal("required CONFIG_PATH enviroment variable is not set")
	}
	cfg, err := ini.Load(configPath)
	if err != nil {
		echo.Logger.Fatal("failed to load config file from path: %s", configPath)
	}

	config := new(Config)
	err = cfg.MapTo(config)
	if err != nil {
		echo.Logger.Fatalf("failed to map config values to config struct: %v", err)
	}

	server := NewServer(echo, config)
	fmt.Println("the dsn: ", config.PostgresDSN)
	fmt.Println("http: ", config.HttpPort)
	fmt.Println("grpc: ", config.GrpcPort)

	db := server.connectToPostgres()
	if db == nil {
		server.Echo.Logger.Fatal("failed to connect to postgres, shutting down...")
	}

	// execute all the migrations that are not on the database yet
	server.migratePostgres()

	// Serve swagger docs
	// swagger/index.html
	server.Echo.GET(swaggerURL, echoSwagger.WrapHandler)

	httpHandler, grpcHandler := setupHandlers(db, echo.Logger)
	server.setupRoutes(httpHandler)

	// Start server
	go server.httpListen()
	go server.grpcListen(grpcHandler)
	server.handleGracefulShutdown()
}

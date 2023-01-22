package main

type Config struct {
	HttpPort                         string `ini:"conn_http_port"`
	GrpcPort                         string `ini:"conn_grpc_port"`
	PostgresDSN                      string `ini:"database_postgres_dsn"`
	NumOfAttemptsToConnectToDatabase int    `ini:"database_num_of_attempts_to_connect"`
}

type Conn struct {
}

type Database struct {
}

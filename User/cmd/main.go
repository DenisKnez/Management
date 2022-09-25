package main

import (
	"context"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

func main() {

	g := gin.Default()

	server := Server{}

	todoServiceConn, err := server.dialTodoService("")
	if err != nil {
		panic(err)
	}

	defer func() {
		err := todoServiceConn.Close()
		if err != nil {
			panic(err)
		}
	}()

	todoServiceClient := server.createTodoClient(todoServiceConn)

	mongoDB, err := server.connectToMongo("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	defer func() {
		err := mongoDB.Disconnect(context.TODO())
		if err != nil {
			panic(err)
		}
	}()

	setupUser(g, mongoDB, todoServiceClient)

	g.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	endless.ListenAndServe(":4242", g)
}

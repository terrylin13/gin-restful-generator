package main

import (
	"github.com/gin-gonic/gin"

	"github.com/terrylin13/gin-restful-generator/internal/example/internal/config"
	"github.com/terrylin13/gin-restful-generator/internal/example/internal/graphql/handle"
	"github.com/terrylin13/gin-restful-generator/internal/example/internal/model"
	"github.com/terrylin13/gin-restful-generator/internal/example/internal/restful"
	"github.com/terrylin13/gin-restful-generator/internal/example/internal/ws"
)

func main() {

	db, err := config.GetDB()
	if err != nil {
		panic("failed to connect database")
	}

	// DB migration
	model.Migrate(db)

	// GraphQL init
	handle.Init()

	e := gin.New()

	// RESTful handle
	restful.Register(e)

	// GraphQL handle
	e.POST("/graphql", handle.GinHandle)

	// Websocket handle
	e.GET("/ws", func(c *gin.Context) {
		ws.WebSocketHandler(c.Writer, c.Request)
	})

	e.Run(":8080")

}

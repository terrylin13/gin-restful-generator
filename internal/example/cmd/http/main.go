package main

import (
	"github.com/gin-gonic/gin"
	// _ "github.com/terrylin13/gin-restful-generator/internal/example/docs/swagger" // import Swagger docs generator
	"github.com/terrylin13/gin-restful-generator/internal/example/internal/api"
	"github.com/terrylin13/gin-restful-generator/internal/example/internal/config"
	"github.com/terrylin13/gin-restful-generator/internal/example/internal/model"
	"github.com/terrylin13/gin-restful-generator/internal/example/internal/ws"
)

// @title Gin RESTful API
// @description The example for Gin RESTful API
// @version 1.0
// @host localhost:8080
// @BasePath /api
func main() {

	db, err := config.GetDB()
	if err != nil {
		panic("failed to connect database")
	}

	model.Migrate(db)

	e := gin.New()

	api.Register(e)

	e.GET("/ws", func(c *gin.Context) {
		ws.WebSocketHandler(c.Writer, c.Request)
	})

	// e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("./docs/swagger/doc.json"))) // Swagger Doc URL

	e.Run(":8080")

}

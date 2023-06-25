package main

import (
	"github.com/gin-gonic/gin"
	"github.com/terrylin13/gin-restful-generator/example/internal/api"
	"github.com/terrylin13/gin-restful-generator/example/internal/config"
	"github.com/terrylin13/gin-restful-generator/example/internal/model"
	"github.com/terrylin13/gin-restful-generator/example/internal/repository"
	"github.com/terrylin13/gin-restful-generator/example/internal/service"
)

func main() {

	db, err := config.GetDB()
	if err != nil {
		panic("failed to connect database")
	}

	model.Migrate(db)

	userRepo := &repository.GormUserRepository{}
	userService := service.NewUserService(userRepo)

	e := gin.Default()

	e.GET("/ws", func(c *gin.Context) {
		api.WebSocketHandler(c.Writer, c.Request)
	})

	g := e.Group("/api")
	{
		g.POST("/user", api.CreateUser(userService))
		g.GET("/user/:id", api.GetUser(userService))
	}

	e.Run(":8080")

}

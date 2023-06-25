package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/terrylin13/gin-restful-generator/internal/example/docs/swagger" // import Swagger docs generator
	"github.com/terrylin13/gin-restful-generator/internal/example/internal/api"
	"github.com/terrylin13/gin-restful-generator/internal/example/internal/config"
	"github.com/terrylin13/gin-restful-generator/internal/example/internal/model"
	"github.com/terrylin13/gin-restful-generator/internal/example/internal/repository"
	"github.com/terrylin13/gin-restful-generator/internal/example/internal/service"
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

	userRepo := &repository.GormUserRepository{}
	userService := service.NewUserService(userRepo)

	e := gin.Default()

	e.GET("/ws", func(c *gin.Context) {
		api.WebSocketHandler(c.Writer, c.Request)
	})

	// url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	// e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("./docs/swagger/doc.json"))) // Swagger Doc URL
	g := e.Group("/api")
	{

		g.POST("/user", api.CreateUser(userService))
		g.GET("/user/:id", api.GetUser(userService))
	}

	e.Run(":8080")

}

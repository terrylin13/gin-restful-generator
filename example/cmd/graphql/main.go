package main

import (
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/terrylin13/gin-restful-generator/example/internal/model"
	"github.com/terrylin13/gin-restful-generator/example/internal/repository"
	"github.com/terrylin13/gin-restful-generator/example/internal/service"
)

func main() {
	userRepo := &repository.GormUserRepository{}
	userService := service.NewUserService(userRepo)

	// Define the User type
	userType := graphql.NewObject(graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	// Define the Query type
	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(string)
					if ok {
						// Fetch user from service
						var user model.User
						err := userService.GetUser(&user, id)
						if err != nil {
							return nil, err
						}
						return user, nil
					}
					return nil, nil
				},
			},
		},
	})

	// Define the Mutation type
	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createUser": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					// name, _ := params.Args["name"].(string)
					user := model.User{
						Name: params.Args["name"].(string),
					}
					// Create new user
					err := userService.CreateUser(&user)
					if err != nil {
						return nil, err
					}
					return user, nil
				},
			},
		},
	})

	// Define the schema with our Query and Mutation types
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	})
	if err != nil {
		panic(err)
	}

	// Create a new GraphQL handler
	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})

	// Initialize Gin and add the handler
	r := gin.Default()
	r.POST("/graphql", func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	})
	r.Run(":8181")
}

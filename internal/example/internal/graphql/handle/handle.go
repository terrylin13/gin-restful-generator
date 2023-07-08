package handle

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/terrylin13/gin-restful-generator/internal/example/internal/graphql/fields"
)

var Schema graphql.Schema
var RootQuery graphql.ObjectConfig
var RootMutation graphql.ObjectConfig
var SchemaConfig graphql.SchemaConfig

func Init() {
	RootQuery = graphql.ObjectConfig{Name: "Query", Fields: fields.AllQueryFields}
	RootMutation = graphql.ObjectConfig{Name: "Mutation", Fields: fields.AllMutationFields}
	SchemaConfig = graphql.SchemaConfig{
		Query:    graphql.NewObject(RootQuery),
		Mutation: graphql.NewObject(RootMutation),
	}
	var schema, err = graphql.NewSchema(SchemaConfig)
	if err != nil {
		log.Fatalf("failed to create schema: %v", err)
	}
	Schema = schema
}

func GinHandle(c *gin.Context) {
	var query struct {
		Query string `json:"query"`
	}

	if err := c.ShouldBindJSON(&query); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	params := graphql.Params{
		Schema:        Schema,
		RequestString: query.Query,
	}

	result := graphql.Do(params)
	if len(result.Errors) > 0 {
		c.JSON(400, gin.H{"errors": result.Errors})
		return
	}

	c.JSON(200, result)
}

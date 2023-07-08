package fields

import (
	"github.com/graphql-go/graphql"
	"github.com/terrylin13/gin-restful-generator/internal/example/internal/config"
	"github.com/terrylin13/gin-restful-generator/internal/example/internal/graphql/schemas"
	"github.com/terrylin13/gin-restful-generator/internal/example/internal/model"
)

// var UserQueryType = graphql.NewObject(graphql.ObjectConfig{
// 	Name:   "Query",
// 	Fields: UserQueryField,
// })

var UserQueryField = &graphql.Field{
	Type:        graphql.NewList(schemas.User),
	Description: "Get User list",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Description: "User ID",
			Type:        graphql.Int,
		},
		"name": &graphql.ArgumentConfig{
			Description: "Search User Name",
			Type:        graphql.String,
		},
		"page": &graphql.ArgumentConfig{
			Description:  "current page",
			Type:         graphql.Int,
			DefaultValue: 1,
		},
		"perPage": &graphql.ArgumentConfig{
			Description:  "page size",
			Type:         graphql.Int,
			DefaultValue: 10,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		var users []model.User

		// 使用 GORM 进行查询
		db, _ := config.GetDB()
		if id, ok := p.Args["id"].(int); ok {
			db = db.Where("id = ?", id)
		}

		if name, ok := p.Args["name"].(string); ok {
			db = db.Where("LIKE ?", "%"+name+"%")
		}

		page := p.Args["page"].(int)
		perPage := p.Args["perPage"].(int)

		offset := (page - 1) * perPage
		limit := perPage

		if err := db.Offset(offset).Limit(limit).Find(&users).Error; err != nil {
			return nil, err
		}
		return users, nil
	},
}

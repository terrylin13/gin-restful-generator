package fields

import (
	"github.com/graphql-go/graphql"
	"github.com/terrylin13/gin-restful-generator/internal/example/internal/config"
	"github.com/terrylin13/gin-restful-generator/internal/example/internal/graphql/schemas"
	"github.com/terrylin13/gin-restful-generator/internal/example/internal/model"
)

// var ScoreItemQueryType = graphql.NewObject(graphql.ObjectConfig{
// 	Name:   "Query",
// 	Fields: ScoreItemField,
// })

var ScoreItemField = &graphql.Field{
	Type:        graphql.NewList(schemas.ScoreItem),
	Description: "Get score item list",
	Args: graphql.FieldConfigArgument{
		"score_item_id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"score_id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"lat": &graphql.ArgumentConfig{
			Type: graphql.Float,
		},
		"lng": &graphql.ArgumentConfig{
			Type: graphql.Float,
		},
		"type": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"score": &graphql.ArgumentConfig{
			Type: graphql.Float,
		},
		"memo": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"images": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"fy_status": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"severity": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"result": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"result_images": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"created_at": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"updated_at": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"page": &graphql.ArgumentConfig{
			Type:         graphql.Int,
			DefaultValue: 1,
		},
		"perPage": &graphql.ArgumentConfig{
			Type:         graphql.Int,
			DefaultValue: 10,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		var items []model.ScoreItem

		// 使用 GORM 进行查询
		db, _ := config.GetDB()

		if id, ok := p.Args["score_item_id"].(string); ok {
			db = db.Where("score_item_id = ?", id)
		}

		if id, ok := p.Args["score_id"].(string); ok {
			db = db.Where("score_id = ?", id)
		}

		if status, ok := p.Args["fy_status"].(int); ok {
			db = db.Where("fy_status = ?", status)
		}

		page := p.Args["page"].(int)
		perPage := p.Args["perPage"].(int)

		offset := (page - 1) * perPage
		limit := perPage
		if err := db.Offset(offset).Limit(limit).Find(&items).Error; err != nil {
			return nil, err
		}
		return items, nil
	},
}

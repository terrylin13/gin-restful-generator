package schemas

import (
	"github.com/graphql-go/graphql"
)

var ScoreItem = graphql.NewObject(graphql.ObjectConfig{
	Name:        "ScoreItem",
	Description: "考核复议项",
	Fields: graphql.Fields{
		"score_item_id": &graphql.Field{
			Type: graphql.String,
		},
		"score_id": &graphql.Field{
			Type: graphql.String,
		},
		"lat": &graphql.Field{
			Type: graphql.Float,
		},
		"lng": &graphql.Field{
			Type: graphql.Float,
		},
		"type": &graphql.Field{
			Type: graphql.String,
		},
		"score": &graphql.Field{
			Type: graphql.Float,
		},
		"memo": &graphql.Field{
			Type: graphql.String,
		},
		"images": &graphql.Field{
			Type: graphql.String,
		},
		"fy_status": &graphql.Field{
			Type: graphql.Int,
		},
		"severity": &graphql.Field{
			Type: graphql.Int,
		},
		"result": &graphql.Field{
			Type: graphql.String,
		},
		"result_images": &graphql.Field{
			Type: graphql.String,
		},
		"created_at": &graphql.Field{
			Type: graphql.Int,
		},
		"updated_at": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

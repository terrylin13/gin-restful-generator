package schemas

import (
	"github.com/graphql-go/graphql"
)

var User = graphql.NewObject(graphql.ObjectConfig{
	Name:        "User",
	Description: "User Schema",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		// "created_at": &graphql.Field{
		// 	Type: graphql.Int,
		// },
		// "updated_at": &graphql.Field{
		// 	Type: graphql.Int,
		// },
		// "articles": &graphql.Field{
		// Type:        graphql.NewList(ScoreItem),
		// Description: "Get socre_items of the score",
		// Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		// 	score := p.Source.(model.Score)
		// 	var items []model.ScoreItem
		// 	db, _ := config.GetDB()
		// 	if err := db.Where("score_id = ?", score.ScoreID).Find(&items).Error; err != nil {
		// 		fmt.Println(err)
		// 		return nil, err
		// 	}

		// 	return items, nil
		// },
		// },
	},
})

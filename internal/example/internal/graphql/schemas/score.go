package schemas

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/terrylin13/gin-restful-generator/internal/example/internal/config"
	"github.com/terrylin13/gin-restful-generator/internal/example/internal/model"
)

var Score = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Score",
	Description: "考核单",
	Fields: graphql.Fields{
		"score_id": &graphql.Field{
			Type: graphql.String,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"pos_type": &graphql.Field{
			Type: graphql.Int,
		},
		"pos_id": &graphql.Field{
			Type: graphql.String,
		},
		"pos_name": &graphql.Field{
			Type: graphql.String,
		},
		"dept_id": &graphql.Field{
			Type: graphql.String,
		},
		"dept_name": &graphql.Field{
			Type: graphql.String,
		},
		"group_dept_id": &graphql.Field{
			Type: graphql.String,
		},
		"group_dept_name": &graphql.Field{
			Type: graphql.String,
		},
		"case_type": &graphql.Field{
			Type: graphql.Int,
		},
		"start_at": &graphql.Field{
			Type: graphql.Int,
		},
		"end_at": &graphql.Field{
			Type: graphql.Int,
		},
		"score": &graphql.Field{
			Type: graphql.Float,
		},
		"flow_score": &graphql.Field{
			Type: graphql.Float,
		},
		"flow_v_star": &graphql.Field{
			Type: graphql.Int,
		},
		"flow_q_star": &graphql.Field{
			Type: graphql.Int,
		},
		"star_score": &graphql.Field{
			Type: graphql.Int,
		},
		"cur_actor_ids": &graphql.Field{
			Type: graphql.String,
		},
		"cur_actors": &graphql.Field{
			Type: graphql.String,
		},
		"status": &graphql.Field{
			Type: graphql.Int,
		},
		"created_by": &graphql.Field{
			Type: graphql.String,
		},
		"created_by_name": &graphql.Field{
			Type: graphql.String,
		},
		"created_at": &graphql.Field{
			Type: graphql.Int,
		},
		"updated_at": &graphql.Field{
			Type: graphql.Int,
		},
		"score_items": &graphql.Field{
			Type:        graphql.NewList(ScoreItem),
			Description: "Get socre_items of the score",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				score := p.Source.(model.Score)
				var items []model.ScoreItem
				db, _ := config.GetDB()
				// 使用 GORM 进行联表查询
				if err := db.Where("score_id = ?", score.ScoreID).Find(&items).Error; err != nil {
					fmt.Println(err)
					return nil, err
				}

				return items, nil
			},
		},
	},
})

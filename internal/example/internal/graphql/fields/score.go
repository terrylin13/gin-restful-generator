package fields

import (
	"github.com/graphql-go/graphql"
	"github.com/terrylin13/gin-restful-generator/internal/example/internal/config"
	"github.com/terrylin13/gin-restful-generator/internal/example/internal/graphql/schemas"
	"github.com/terrylin13/gin-restful-generator/internal/example/internal/model"
)

// var ScoreQueryType = graphql.NewObject(graphql.ObjectConfig{
// 	Name:   "Query",
// 	Fields: ScoreQueryField,
// })

var ScoreQueryField = &graphql.Field{
	Type:        graphql.NewList(schemas.Score),
	Description: "Get score list",
	Args: graphql.FieldConfigArgument{
		"score_id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"title": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"pos_type": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"pos_id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"pos_name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"dept_id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"dept_name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"group_dept_id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"group_dept_name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"case_type": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"start_at": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"end_at": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"score": &graphql.ArgumentConfig{
			Type: graphql.Float,
		},
		"flow_score": &graphql.ArgumentConfig{
			Type: graphql.Float,
		},
		"flow_v_star": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"flow_q_star": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"star_score": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"cur_actor_ids": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"cur_actors": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"status": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"status2": &graphql.ArgumentConfig{
			Type: graphql.NewList(graphql.Int),
		},
		"created_by": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"created_by_name": &graphql.ArgumentConfig{
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
		var scores []model.Score

		// 使用 GORM 进行查询
		db, _ := config.GetDB()
		if id, ok := p.Args["score_id"].(string); ok {
			db = db.Where("score_id = ?", id)
		}

		if status, ok := p.Args["status"].(int); ok {
			db = db.Where("status = ?", status)
		}

		if status2, ok := p.Args["status2"].([]interface{}); ok {
			db = db.Where("status IN ?", status2)
		}

		page := p.Args["page"].(int)
		perPage := p.Args["perPage"].(int)

		offset := (page - 1) * perPage
		limit := perPage
		if err := db.Offset(offset).Limit(limit).Find(&scores).Error; err != nil {
			return nil, err
		}
		return scores, nil
	},
}

var CreateScore = &graphql.Field{
	Type: schemas.Score,
	Args: graphql.FieldConfigArgument{
		"score_id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"title": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"pos_type": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"pos_id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"pos_name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"dept_id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"dept_name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"group_dept_id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"group_dept_name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"case_type": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"start_at": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"end_at": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"score": &graphql.ArgumentConfig{
			Type: graphql.Float,
		},
		"flow_score": &graphql.ArgumentConfig{
			Type: graphql.Float,
		},
		"flow_v_star": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"flow_q_star": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"star_score": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"created_by": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"created_by_name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		title, _ := p.Args["title"].(string)
		posType, _ := p.Args["pos_type"].(int)
		posID, _ := p.Args["pos_id"].(string)
		s := model.Score{
			Title:   title,
			PosType: uint8(posType),
			PosID:   posID,
		}
		// 创建并存储到数据库

		return s, nil
	},
}

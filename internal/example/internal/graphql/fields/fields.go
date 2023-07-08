package fields

import "github.com/graphql-go/graphql"

var AllQueryFields = graphql.Fields{
	"score_item": ScoreItemField,
	"score":      ScoreQueryField,
	"user":       UserQueryField,
}

var AllMutationFields = graphql.Fields{
	"createScore": CreateScore,
}

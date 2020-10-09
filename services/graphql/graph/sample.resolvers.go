package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/takuya911/golang-api_sample/services/graphql/graph/generated"
	"github.com/takuya911/golang-api_sample/services/graphql/graph/model"
	sample "github.com/takuya911/golang-api_sample/services/graphql/proto"
)

func (r *mutationResolver) CreateRecode(ctx context.Context, input sample.CreateRecodeForm) (*sample.Recode, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetRecodeByID(ctx context.Context, input model.GetRecodeForm) (*sample.Recode, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *resolver }
type queryResolver struct{ *resolver }

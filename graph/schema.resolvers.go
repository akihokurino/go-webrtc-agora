package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"canvas-server/graph/generated"
	"canvas-server/graph/model"
	"context"
	"fmt"
)

func (r *queryResolver) RtcToken(ctx context.Context) (*model.RTCToken, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
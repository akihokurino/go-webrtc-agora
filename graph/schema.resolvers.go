package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"go-webrtc/graph/generated"
	"go-webrtc/graph/model"
)

func (r *queryResolver) AgoraToken(ctx context.Context, channelName string) (*model.AgoraToken, error) {
	uid := r.contextProvider.MustAuthUID(ctx)
	token, err := r.agoraClient.GetRTCToken(uid, channelName)
	if err != nil {
		return nil, err
	}

	return &model.AgoraToken{
		Token: token,
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

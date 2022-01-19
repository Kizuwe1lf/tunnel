package api

import (
	"context"
	pb "tunnel/proto/gen/v1"
)

type ServiceImplementation struct {
	pb.UnimplementedTunnelServiceServer
}

func (s *ServiceImplementation) PostTweet(ctx context.Context, request pb.PostTweetRequest) (pb.PostTweetResponse, error) {
	return pb.PostTweetResponse{}, nil
}

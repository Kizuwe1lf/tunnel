package api

import (
	"context"
	"tunnel/client"
	pb "tunnel/proto/gen/v1"
)

type Service struct {
	pb.UnimplementedTunnelServiceServer
	twitter client.Twitter
}

func (s *Service) PostTweet(ctx context.Context, request *pb.PostTweetRequest) (*pb.PostTweetResponse, error) {
	mediaID, err := s.twitter.UploadMedia(request.Image)

	if err != nil {
		return &pb.PostTweetResponse{}, err
	}

	err = s.twitter.PostMediaTweet(mediaID)

	if err != nil {
		return &pb.PostTweetResponse{}, err
	}

	return &pb.PostTweetResponse{}, nil
}

func NewService(token, tokenSecret, consumerKey, consumerSecret string) *Service {
	return &Service{
		twitter: client.NewClient(token, tokenSecret, consumerKey, consumerSecret),
	}
}

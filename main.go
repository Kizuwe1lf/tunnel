package main

import (
	"github.com/golobby/dotenv"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"tunnel/api"
	pb "tunnel/proto/gen/v1"
)

type Credentials struct {
	Token          string `env:"TOKEN"`
	TokenSecret    string `env:"TOKEN_SECRET"`
	ConsumerKey    string `env:"CONSUMER_KEY"`
	ConsumerSecret string `env:"CONSUMER_SECRET"`
}

func main() {
	credentials := Credentials{}
	file, err := os.Open(".env")
	err = dotenv.NewDecoder(file).Decode(&credentials)

	lis, err := net.Listen("tcp", ":1333")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()
	
	s := grpc.NewServer()
	service := api.NewService(credentials.Token, credentials.TokenSecret, credentials.ConsumerKey, credentials.ConsumerSecret)
	pb.RegisterTunnelServiceServer(s, service)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

package main

import (
	"context"
	"net"

	pb "github.com/dink10/lessons/grpc/api"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	port = ":3001"
)

type server struct {
	pb.UnimplementedChatServiceServer
}

func (s *server) SayHello(_ context.Context, in *pb.Message) (*pb.Message, error) {
	logrus.Infof("Received greeting from client: %s", in.Body)

	return &pb.Message{Body: "Hey"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		logrus.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterChatServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		logrus.Fatal(err)
	}
}

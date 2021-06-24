package main

import (
	"context"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	pb "github.com/dink10/lessons/grpc/api"
)

const (
	address = "localhost:3001"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		logrus.Fatal(err)
	}
	defer func() {
		_ = conn.Close()
	}()

	client := pb.NewChatServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	message := pb.Message{Body: "Hello"}

	resp, err := client.SayHello(ctx, &message)
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}

	logrus.Infof("Greetings from server: %s", resp.Body)
}

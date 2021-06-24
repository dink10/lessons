package main

import (
	"database/sql"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"

	_ "github.com/lib/pq"

	pb "github.com/dink10/lessons/grpc_users_service/api"
	"github.com/dink10/lessons/grpc_users_service/internal/service"
	"github.com/sirupsen/logrus"
)

const (
	port = ":3001"
)

func main() {
	logrus.Info("GRPC service started")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	connUrl := "postgres://admin:admin@localhost:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connUrl)
	if err != nil {
		logrus.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		logrus.Fatal(err)
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		logrus.Fatal(err)
	}

	srv := service.NewGRPCServer(db)

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, srv)

	go func() {
		if err := s.Serve(lis); err != nil {
			logrus.Fatal(err)
		}
	}()

	<-stop

	logrus.Info("Gracefully stopped")
}

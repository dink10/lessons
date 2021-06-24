package service

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/dink10/lessons/grpc_users_service/internal/entity"

	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/dink10/lessons/grpc_users_service/api"
)

type grpcServer struct {
	pb.UnimplementedUserServiceServer
	db *sql.DB
}

func NewGRPCServer(db *sql.DB) pb.UserServiceServer {
	return &grpcServer{db: db}
}

func (s *grpcServer) CreateUser(ctx context.Context, in *pb.User) (*pb.User, error) {

	item := pbToUser(in)

	query := `INSERT INTO users(name, age, email) VALUES($1, $2, $3);`

	_, err := s.db.ExecContext(ctx, query, item.Name, item.Age, item.Email)
	if err != nil {
		return nil, err
	}

	return entityToModel(item), nil
}

func (s *grpcServer) UpdateUser(ctx context.Context, in *pb.User) (*pb.User, error) {
	item := pbToUser(in)

	query := "UPDATE users SET name=?, age=?, email=? WHERE id=$1;"

	_, err := s.db.ExecContext(ctx, query, item.Name, item.Age, item.Email, item.ID)
	if err != nil {
		return nil, err
	}

	return entityToModel(item), nil
}

func (s *grpcServer) GetUsers(ctx context.Context, _ *emptypb.Empty) (*pb.Users, error) {

	query := "SELECT id, name, age, email FROM users"

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var items entity.Users
	for rows.Next() {
		u, errScan := userScan(rows)
		if errScan != nil {
			return nil, errScan
		}

		items = append(items, u)
	}

	return entitiesToModels(items), nil
}

func (s *grpcServer) DeleteUser(ctx context.Context, in *pb.Identifier) (*emptypb.Empty, error) {

	id := int(in.Id)

	query := "DELETE FROM users WHERE id=$1;"
	res, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		return nil, err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rows == 0 {
		return nil, fmt.Errorf("user with id %d not found", id)
	}

	return &emptypb.Empty{}, nil
}

// Scanner provides scan interface
type Scanner interface {
	Scan(dest ...interface{}) error
}

// scan helper
func userScan(rows Scanner) (*entity.User, error) {
	var (
		id, age     int
		name, email string
	)

	err := rows.Scan(
		&id,
		&name,
		&age,
		&email,
	)
	if err != nil {
		return nil, err
	}

	return &entity.User{
		ID:    id,
		Name:  name,
		Age:   age,
		Email: email,
	}, err
}

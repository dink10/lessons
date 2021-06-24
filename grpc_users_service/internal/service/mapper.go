package service

import (
	pb "github.com/dink10/lessons/grpc_users_service/api"

	"github.com/dink10/lessons/grpc_users_service/internal/entity"
)

func pbToUser(in *pb.User) *entity.User {
	return &entity.User{
		Name:  in.Name,
		Age:   int(in.Age),
		Email: in.Email,
	}
}

func entityToModel(item *entity.User) *pb.User {
	return &pb.User{
		Id:    int64(item.ID),
		Name:  item.Name,
		Age:   int64(item.Age),
		Email: item.Email,
	}
}

func entitiesToModels(items entity.Users) *pb.Users {
	var res pb.Users
	for _, i := range items {
		res.Users = append(res.Users, &pb.User{
			Id:    int64(i.ID),
			Name:  i.Name,
			Age:   int64(i.Age),
			Email: i.Email,
		})
	}

	return &res
}

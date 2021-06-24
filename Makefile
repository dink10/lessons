gen:
	protoc --go_out=plugins=grpc:./ --go_opt=module=github.com/dink10/lessons ./grpc/api/chat.proto
	protoc --go_out=plugins=grpc:./ --go_opt=module=github.com/dink10/lessons ./grpc_users_service/api/user.proto
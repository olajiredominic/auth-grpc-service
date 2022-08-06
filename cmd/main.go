package main

import (
	"log"
	"net"

	"cityhotels.com/backend-auth/pkg/config"
	"cityhotels.com/backend-auth/pkg/pb"
	"cityhotels.com/backend-auth/pkg/routes"

	"google.golang.org/grpc"
)

func main() {

	config, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	lis, err := net.Listen("tcp", config.Port)
	if err != nil {
		log.Fatalln("Failed to serve:", err)
	}

	handler := routes.Init(config.DBUrl)
	h := routes.New(handler.DB)

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &h)
	// pb.RegisterRoomServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}

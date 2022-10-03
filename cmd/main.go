package main

import (
	"fmt"
	"log"
	"net"

	"github.com/lerryjay/auth-grpc-service/pkg/config"
	"github.com/lerryjay/auth-grpc-service/pkg/pb"
	"github.com/lerryjay/auth-grpc-service/pkg/routes"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

	dbUrl := fmt.Sprintf("postgres://%s:%s@%s", config.DBUSER, config.DBPWD, config.DBURL)
	fmt.Printf("Database Url", dbUrl)
	handler := routes.Init(dbUrl)
	h := routes.New(handler.DB)

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &h)
	pb.RegisterAuthServiceServer(grpcServer, &h)
	pb.RegisterTestServiceServer(grpcServer, &h)
	pb.RegisterHealthServer(grpcServer, &h)

	reflection.Register(grpcServer)
	// pb.RegisterRoomServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}

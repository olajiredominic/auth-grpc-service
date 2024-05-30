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
	log.Println("Database Url", dbUrl)
	handler := routes.Init(dbUrl, config.CLIENT_ID, config.SECRET_KEY, config.TOKEN_URL, config.QOREID_BASE_URL, config.VNIN_URL, config.NIN_URL)
	h := routes.Handler{
		DB:            handler.DB,
		ClientID:      config.CLIENT_ID,
		SecretKey:     config.SECRET_KEY,
		TokenURL:      config.TOKEN_URL,
		QoreidBaseURL: config.QOREID_BASE_URL,
		VNINURL:       config.VNIN_URL,
		NINURL:        config.NIN_URL,
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &h)
	pb.RegisterAuthServiceServer(grpcServer, &h)
	pb.RegisterTestServiceServer(grpcServer, &h)
	pb.RegisterVerificationServiceServer(grpcServer, &h)
	pb.RegisterHealthServer(grpcServer, &h)

	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}

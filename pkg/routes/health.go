package routes

import (
	"log"

	"cityhotels.com/backend-auth/pkg/pb"
	"golang.org/x/net/context"
)

func (s *Handler) Check(ctx context.Context, req *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	log.Println("Serving the Check request for health check")
	return &pb.HealthCheckResponse{
		Status: pb.HealthCheckResponse_SERVING,
	}, nil
}

func (s *Handler) Watch(ctx context.Context, req *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	log.Println("Serving the Watch request for health check")
	return &pb.HealthCheckResponse{
		Status: pb.HealthCheckResponse_SERVING,
	}, nil
}

package routes

import (
	"github.com/City-Hotels/ch-backend-auth/pkg/pb"
	"golang.org/x/net/context"
)

func (h *Handler) SayingHello(ctx context.Context, req *pb.SayingHelloRequest) (*pb.SayingHelloResponse, error) {

	return &pb.SayingHelloResponse{Response: "Hello " + req.Name}, nil
}

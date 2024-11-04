package controller

import (
	pb "github.com/nurzzaat/gRPC-example/auth/proto"
	"github.com/nurzzaat/gRPC-example/auth/repository"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	pb.UnimplementedAuthServiceServer
	UserRepository repository.AuthRepository
}

func NewGRPCHandler(grpcServer *grpc.Server, repo repository.AuthRepository) {
	handler := &grpcHandler{
		UserRepository: repo,
	}
	pb.RegisterAuthServiceServer(grpcServer, handler)
}

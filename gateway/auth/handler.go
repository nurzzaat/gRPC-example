package auth

import (
	pb "github.com/nurzzaat/gRPC-example/auth/proto"

	log "github.com/sirupsen/logrus"

	"google.golang.org/grpc"
)

type AuthController struct {
	Client pb.AuthServiceClient
}

func NewAuthController(client pb.AuthServiceClient)*AuthController{
	return &AuthController{client}
}
const (
	authServiceAddr = "localhost:8000"
)

func NewAuthClient() (pb.AuthServiceClient, error) {
	conn, err := grpc.Dial(authServiceAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Close unexpectedly with error:", err.Error())
	}
	//defer conn.Close()

	log.Println("Dialing auth service at:", authServiceAddr)
	client := pb.NewAuthServiceClient(conn)
	return client, nil
}

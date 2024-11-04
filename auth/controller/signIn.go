package controller

import (
	"context"
	"fmt"

	"github.com/nurzzaat/gRPC-example/auth/pkg"
	pb "github.com/nurzzaat/gRPC-example/auth/proto"
	"github.com/nurzzaat/gRPC-example/auth/tokenutil"
	log "github.com/sirupsen/logrus"
	//"golang.org/x/crypto/bcrypt"
)

func (s *grpcHandler) SignIn(ctx context.Context, request *pb.SignInRequest) (*pb.SignInResponse, error) {
	logFields := log.Fields{
		"requestType": "POST",
		"endpoint":    "/auth/sign-in",
	}

	fmt.Println("It's come here!!!")
	if request.Email == "" || request.Password == "" {
		log.WithFields(logFields).Error("Empty values are declared:", request)
		return nil, fmt.Errorf("104")
	}

	user, err := s.GetUserByEmail(ctx, &pb.UserEmail{Email: request.Email})
	if err != nil {
		log.WithFields(logFields).Error("Get admin userinfo error:", err.Error())
		return nil, fmt.Errorf("...")
	}
	// if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
	// 	log.WithFields(logFields).Error("Hash and Password are different")
	// 	return nil, fmt.Errorf("102")
	// }

	accessToken, err := tokenutil.CreateAccessToken(ctx, uint(user.Id), pkg.ACCESS_TOKEN_SECRET, pkg.ACCESS_TOKEN_EXPIRY_HOUR, pkg.NewRedisConnection())
	if err != nil {
		log.WithFields(logFields).Error(err.Error())
		return nil, fmt.Errorf("102")
	}
	return &pb.SignInResponse{Token: accessToken}, nil
}

func (s *grpcHandler) GetUserByEmail(ctx context.Context, userRequest *pb.UserEmail) (*pb.UserResponse, error) {
	user, _ := s.UserRepository.GetUserByEmail(ctx, userRequest.Email)
	return &pb.UserResponse{Id: int32(user.ID), Password: user.Password, Email: user.Email}, nil
}

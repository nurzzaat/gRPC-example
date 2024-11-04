package controller

import (
	"context"
	"fmt"

	emailverifier "github.com/AfterShip/email-verifier"
	"github.com/nurzzaat/gRPC-example/auth/pkg"
	pb "github.com/nurzzaat/gRPC-example/auth/proto"
	"github.com/nurzzaat/gRPC-example/auth/repository"
	"github.com/nurzzaat/gRPC-example/auth/tokenutil"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type Signup struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var (
	verifier = emailverifier.NewVerifier()
)

func (sc *grpcHandler) SignUp(ctx context.Context, request *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	logFields := log.Fields{
		"requestType": "POST",
		"endpoint":    "/auth/sign-up",
	}

	verifier = verifier.EnableSMTPCheck()
	verifier = verifier.EnableDomainSuggest()

	if request.Email == "" || request.Password == "" {
		log.WithFields(logFields).Error("Empty values are declared:", request)
		return nil, fmt.Errorf("104")
	}

	user, _ := sc.GetUserByEmail(ctx, &pb.UserEmail{Email: request.Email})
	if user.Id > 0 {
		log.WithFields(logFields).Error("User already exists")
		return nil, fmt.Errorf("104")
	}

	ret, _ := verifier.Verify(request.Email)
	if !ret.Syntax.Valid {
		log.WithFields(logFields).Error("email address syntax is invalid")
		return nil, fmt.Errorf("104")
	}

	if len(request.Password) < 6 {
		log.WithFields(logFields).Error("Password length must be longer than 6 character")
		return nil, fmt.Errorf("104")
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		log.WithFields(logFields).Error("Error generating hash password", err.Error())
		return nil, fmt.Errorf("104")
	}

	newUser := repository.User{
		Email:    request.Email,
		Password: string(encryptedPassword),
	}
	err = sc.UserRepository.CreateUser(ctx, newUser)
	if err != nil {
		log.WithFields(logFields).Error("Create user error:", err.Error())
		return nil, fmt.Errorf("104")
	}
	accessToken, err := tokenutil.CreateAccessToken(ctx, uint(user.Id), pkg.ACCESS_TOKEN_SECRET, pkg.ACCESS_TOKEN_EXPIRY_HOUR, pkg.NewRedisConnection())
	if err != nil {
		log.WithFields(logFields).Error(err.Error())
		return nil, fmt.Errorf("102")
	}
	return &pb.SignUpResponse{Token: accessToken}, nil
}

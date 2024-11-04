package main

import (
	"log"
	"net"

	"github.com/nurzzaat/gRPC-example/auth/controller"
	"github.com/nurzzaat/gRPC-example/auth/pkg"
	"github.com/nurzzaat/gRPC-example/auth/repository"
	"google.golang.org/grpc"
)

//how to implement microservice in golang

var (
	grpcAddr = "localhost:8000"
)

func main() {
	grpcServer := grpc.NewServer()

	pqlDB, err := pkg.NewPgxConn()
	if err != nil {
		log.Fatal("Error connection to db:", err.Error())
	}

	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatal("")
	}
	defer l.Close()

	repo := repository.NewUserRepository(pqlDB)
	controller.NewGRPCHandler(grpcServer, repo)

	log.Println("GRPC server started at", grpcAddr)
	if err := grpcServer.Serve(l); err != nil {
		log.Fatal(err.Error())
	}
}

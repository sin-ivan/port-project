package main

import (
	"log"
	"net"

	deliveryGrpc "github.com/port-project/domain-service/port/delivery/grpc"
	portRepo "github.com/port-project/domain-service/port/repository"
	portUsecase "github.com/port-project/domain-service/port/usecase"
	"google.golang.org/grpc"
)

const (
	serverPort = ":8081"
)

func main() {
	repo := portRepo.NewMapPortRepository()
	usecase := portUsecase.NewPortUsecase(repo)

	listener, err := net.Listen("tcp", serverPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	deliveryGrpc.NewPortServerGrpc(server, usecase)
	log.Println("Starting service...")
	err = server.Serve(listener)
	if err != nil {
		log.Println("Unexpected error:", err)
	}
}

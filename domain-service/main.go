package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

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

	waitForShutdown(server)
}

func waitForShutdown(srv *grpc.Server) {

	go func() {
		interruptChan := make(chan os.Signal, 1)
		signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

		// Block until we receive our signal.
		<-interruptChan

		srv.GracefulStop()

		log.Println("Shutting down")
		println("Shut down server")
		os.Exit(0)
	}()
}

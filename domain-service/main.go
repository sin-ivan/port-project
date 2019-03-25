package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	handler "github.com/port-project/domain-service/port/delivery/grpc/handler"
	repository "github.com/port-project/domain-service/port/repository"
	usecase "github.com/port-project/domain-service/port/usecase"
	"google.golang.org/grpc"
)

const (
	serverPort = ":8081"
)

func main() {
	repo := repository.NewMapPortRepository()
	ucase := usecase.NewPortUsecase(repo)

	listener, err := net.Listen("tcp", serverPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("Starting service...")

	serv := grpc.NewServer()
	handler.NewPortServerGrpc(serv, ucase)
	err = serv.Serve(listener)
	if err != nil {
		log.Println("Unexpected error:", err)
	}

	waitForShutdown(serv)
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

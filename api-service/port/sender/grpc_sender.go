package sender

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	models "github.com/sin-ivan/port-project/api-service/port"
	"github.com/sin-ivan/port-proto"
	"google.golang.org/grpc"
)

type sender struct {
}

func address() string {
	url := os.Getenv("DOMAIN_URL")
	if len(url) == 0 {
		url = "localhost:8081"
	}
	return url
}

// transformToPortRPC is used to map model Port to gRPC port objects
func (s *sender) transformPortDataToRPC(port *models.Port) *proto_grpc.Port {
	if port == nil {
		return nil
	}

	response := &proto_grpc.Port{
		ID:          port.ID,
		Name:        port.Name,
		City:        port.City,
		Country:     port.Country,
		Coordinates: port.Coordinates,
		Alias:       port.Alias,
		Regions:     port.Regions,
		Province:    port.Province,
		Timezone:    port.Timezone,
		Unlocs:      port.Unlocs,
		Code:        port.Code,
	}
	return response
}

func (s *sender) transformPortRPCToData(port *proto_grpc.Port) *models.Port {
	if port == nil {
		return nil
	}

	response := &models.Port{
		ID:          port.ID,
		Name:        port.Name,
		City:        port.City,
		Country:     port.Country,
		Coordinates: port.Coordinates,
		Alias:       port.Alias,
		Regions:     port.Regions,
		Province:    port.Province,
		Timezone:    port.Timezone,
		Unlocs:      port.Unlocs,
		Code:        port.Code,
	}
	return response
}

var _connection *grpc.ClientConn

func grpcConnection() *grpc.ClientConn {

	// Contact the server and print out its response.
	// Set up a connection to the server.

	if _connection != nil {
		return _connection
	}

	var err error
	connection, err := grpc.Dial(address(), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	_connection = connection

	return connection
}

func waitForShutdown() {

	go func() {
		interruptChan := make(chan os.Signal, 1)
		signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

		// Block until we receive our signal.
		<-interruptChan

		defer _connection.Close()

		log.Println("Closing gRPC connection")
		os.Exit(0)
	}()
}

// NewGrpcSender create new instance of gRPC data sender
func NewGrpcSender() PortSender {
	waitForShutdown()
	return &sender{}
}

// StorePort is used to store port to DB using gRPC
func (s *sender) StorePort(port *models.Port) {

	client := proto_grpc.NewPortHandlerClient(grpcConnection())
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	p := s.transformPortDataToRPC(port)

	// Perform request to store port information
	_, err := client.Store(ctx, p)
	if err != nil {
		log.Println("Could not sent port info:", err)
	}
}

func (s *sender) GetAll() []*models.Port {
	log.Println("Getting port list")

	client := proto_grpc.NewPortHandlerClient(grpcConnection())
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Request all the available items
	r, err := client.GetAll(ctx, &proto_grpc.EmptyRequest{})
	if err != nil {
		log.Fatalf("could not get port info: %v", err)
	}

	// Transform gRPC item list to model objects
	items := make([]*models.Port, len(r.Ports))
	for i, p := range r.Ports {
		port := s.transformPortRPCToData(p)
		items[i] = port
	}

	return items
}

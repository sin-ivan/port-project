package sender

import (
	"context"
	"log"
	"os"
	"time"

	models "github.com/port-project/api-service/port"
	"github.com/port-project/api-service/port/delivery/grpc/port_grpc"
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
func (s *sender) transformPortDataToRPC(port *models.Port) *port_grpc.Port {
	if port == nil {
		return nil
	}

	response := &port_grpc.Port{
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

func (s *sender) transformPortRPCToData(port *port_grpc.Port) *models.Port {
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

// NewGrpcSender create new instance of gRPC data sender
func NewGrpcSender() PortSender {
	return &sender{}
}

// StorePort is used to store port to DB using gRPC
func (s *sender) StorePort(port *models.Port) {

	// Contact the server and print out its response.
	// Set up a connection to the server.
	var err error
	conn, err := grpc.Dial(address(), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := port_grpc.NewPortHandlerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	p := s.transformPortDataToRPC(port)

	_, err = client.Store(ctx, p)
	if err != nil {
		log.Println("Could not sent port info:", err)
	}
}

func (s *sender) GetAll() []*models.Port {
	log.Println("Getting port list")

	// Contact the server and print out its response.
	// Set up a connection to the server.
	conn, err := grpc.Dial(address(), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := port_grpc.NewPortHandlerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	// Request all the available items
	r, err := client.GetAll(ctx, &port_grpc.EmptyRequest{})
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

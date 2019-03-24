package grpc

import (
	"context"
	"log"

	models "github.com/port-project/domain-service/port"
	"github.com/port-project/domain-service/port/delivery/grpc/port_grpc"
	usecase "github.com/port-project/domain-service/port/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	portUsecase usecase.PortUsecase
}

// transformToPortRPC is used to map model Port to gRPC port objects
func (s *server) transformPortDataToRPC(port *models.Port) *port_grpc.Port {
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

func (s *server) transformPortRPCToData(port *port_grpc.Port) *models.Port {
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

// NewPortServerGrpc is used to create server instance
func NewPortServerGrpc(grpcServer *grpc.Server, portUsecase usecase.PortUsecase) {
	portServer := &server{
		portUsecase: portUsecase,
	}

	port_grpc.RegisterPortHandlerServer(grpcServer, portServer)
	reflection.Register(grpcServer)
}

func (s *server) GetPort(ctx context.Context, in *port_grpc.SingleRequest) (*port_grpc.Port, error) {
	port, err := s.portUsecase.GetByID(in.Id)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	if port == nil {
		return nil, models.NOT_FOUND_ERROR
	}

	response := s.transformPortDataToRPC(port)
	return response, nil
}

func (s *server) GetAll(ctx context.Context, in *port_grpc.EmptyRequest) (*port_grpc.ListPort, error) {
	ports, err := s.portUsecase.GetAll()
	if err != nil {
		log.Println("Failed getting item list")
	}

	items := make([]*port_grpc.Port, len(ports))
	for i, p := range ports {
		port := s.transformPortDataToRPC(p)
		items[i] = port
	}

	response := &port_grpc.ListPort{
		Ports: items,
	}
	return response, nil
}

func (s *server) Store(ctx context.Context, in *port_grpc.Port) (*port_grpc.Port, error) {
	port := s.transformPortRPCToData(in)
	p, err := s.portUsecase.Store(port)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	response := s.transformPortDataToRPC(p)
	return response, nil
}

func (s *server) Update(ctx context.Context, in *port_grpc.Port) (*port_grpc.Port, error) {
	port := s.transformPortRPCToData(in)
	p, err := s.portUsecase.Store(port)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	response := s.transformPortDataToRPC(p)
	return response, nil
}

func (s *server) Delete(ctx context.Context, in *port_grpc.SingleRequest) (*port_grpc.DeleteResponse, error) {

	removed, err := s.portUsecase.Delete(in.Id)
	if err != nil {
		log.Println("Failed deleting item")
	}

	response := &port_grpc.DeleteResponse{
		Status: "Failed to delete",
	}

	if removed {
		response.Status = "Successfully deleted"
	}

	return response, nil
}

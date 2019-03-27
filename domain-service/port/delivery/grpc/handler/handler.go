package grpc

import (
	"context"
	"log"

	models "github.com/sin-ivan/port-project/domain-service/port"
	usecase "github.com/sin-ivan/port-project/domain-service/port/usecase"
	"github.com/sin-ivan/port-proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	portDeletedCode       = 0
	portWasNotDeletedCode = -1000
)

// PortGrpcHandler is used for handling gRPC callbacks
type PortGrpcHandler interface {
	GetPort(ctx context.Context, in *proto_grpc.SingleRequest) (*proto_grpc.Port, error)
	GetAll(ctx context.Context, in *proto_grpc.EmptyRequest) (*proto_grpc.ListPort, error)
	Store(ctx context.Context, in *proto_grpc.Port) (*proto_grpc.Port, error)
	Update(ctx context.Context, in *proto_grpc.Port) (*proto_grpc.Port, error)
	Delete(ctx context.Context, in *proto_grpc.SingleRequest) (*proto_grpc.DeleteResponse, error)
}

type server struct {
	portUsecase usecase.PortUsecase
}

// NewPortServerGrpc is used to create server instance
func NewPortServerGrpc(grpcServer *grpc.Server, portUsecase usecase.PortUsecase) PortGrpcHandler {
	portServer := &server{
		portUsecase: portUsecase,
	}

	if grpcServer != nil {
		proto_grpc.RegisterPortHandlerServer(grpcServer, portServer)
		reflection.Register(grpcServer)
	}
	return portServer
}

func (s *server) GetPort(ctx context.Context, in *proto_grpc.SingleRequest) (*proto_grpc.Port, error) {
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

func (s *server) GetAll(ctx context.Context, in *proto_grpc.EmptyRequest) (*proto_grpc.ListPort, error) {
	ports, err := s.portUsecase.GetAll()
	if err != nil {
		log.Println("Failed getting item list")
	}

	items := make([]*proto_grpc.Port, len(ports))
	for i, p := range ports {
		port := s.transformPortDataToRPC(p)
		items[i] = port
	}

	response := &proto_grpc.ListPort{
		Ports: items,
	}
	return response, nil
}

func (s *server) Store(ctx context.Context, in *proto_grpc.Port) (*proto_grpc.Port, error) {
	port := s.transformPortRPCToData(in)
	p, err := s.portUsecase.Store(port)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	response := s.transformPortDataToRPC(p)
	return response, nil
}

func (s *server) Update(ctx context.Context, in *proto_grpc.Port) (*proto_grpc.Port, error) {
	port := s.transformPortRPCToData(in)
	p, err := s.portUsecase.Update(port)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	response := s.transformPortDataToRPC(p)
	return response, nil
}

func (s *server) Delete(ctx context.Context, in *proto_grpc.SingleRequest) (*proto_grpc.DeleteResponse, error) {

	removed, err := s.portUsecase.Delete(in.Id)
	if err != nil {
		log.Println("Failed deleting item")
	}

	response := &proto_grpc.DeleteResponse{
		Status: "Failed to delete",
		Code:   portWasNotDeletedCode,
	}

	if removed {
		response.Status = "Successfully deleted"
		response.Code = portDeletedCode
	}

	return response, err
}

// transformToPortRPC is used to map model Port to gRPC port objects
func (s *server) transformPortDataToRPC(port *models.Port) *proto_grpc.Port {
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

func (s *server) transformPortRPCToData(port *proto_grpc.Port) *models.Port {
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

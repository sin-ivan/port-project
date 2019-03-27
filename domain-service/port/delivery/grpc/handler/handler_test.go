package grpc_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	models "github.com/sin-ivan/port-project/domain-service/port"
	handler "github.com/sin-ivan/port-project/domain-service/port/delivery/grpc/handler"
	"github.com/sin-ivan/port-project/domain-service/port/usecase/mocks"
	"github.com/sin-ivan/port-proto"
	"google.golang.org/grpc"
)

func TestGetByIDWithItems(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mock_usecase.NewMockPortUsecase(ctrl)

	identifier := "ABCDE"

	port := &models.Port{
		ID: identifier,
	}

	mockUsecase.EXPECT().GetByID(port.ID).Return(port, nil)

	var server *grpc.Server
	ctx := context.Background()
	req := &proto_grpc.SingleRequest{
		Id: identifier,
	}

	s := handler.NewPortServerGrpc(server, mockUsecase)
	p, err := s.GetPort(ctx, req)

	if err != nil || p.ID != port.ID {
		t.Error("item should be fetched in filled repository")
	}
}

func TestGetByIDNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mock_usecase.NewMockPortUsecase(ctrl)

	identifier := "ABCDE"

	port := &models.Port{
		ID: identifier,
	}

	mockUsecase.EXPECT().GetByID(port.ID).Return(nil, nil)

	var server *grpc.Server
	ctx := context.Background()
	req := &proto_grpc.SingleRequest{
		Id: identifier,
	}

	s := handler.NewPortServerGrpc(server, mockUsecase)
	p, err := s.GetPort(ctx, req)

	if err != models.NOT_FOUND_ERROR || p != nil {
		t.Error("item should not be fetched in empty repository")
	}
}

func TestGetByIDNoItems(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mock_usecase.NewMockPortUsecase(ctrl)

	identifier := "ABCDE"

	err := models.NOT_FOUND_ERROR
	mockUsecase.EXPECT().GetByID(identifier).Return(nil, err)

	var server *grpc.Server
	ctx := context.Background()
	req := &proto_grpc.SingleRequest{
		Id: identifier,
	}

	s := handler.NewPortServerGrpc(server, mockUsecase)
	p, err := s.GetPort(ctx, req)

	if err != models.NOT_FOUND_ERROR || p != nil {
		t.Error("item should not be fetched in empty repository")
	}
}

func TestGetAllNoItems(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mock_usecase.NewMockPortUsecase(ctrl)

	mockUsecase.EXPECT().GetAll().Return(nil, nil)

	var server *grpc.Server
	ctx := context.Background()
	req := &proto_grpc.EmptyRequest{}

	s := handler.NewPortServerGrpc(server, mockUsecase)
	p, err := s.GetAll(ctx, req)

	if err != nil || len(p.Ports) != 0 {
		t.Error("item should not be fetched in empty repository")
	}
}

func TestGetAllFailedFetching(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mock_usecase.NewMockPortUsecase(ctrl)

	err := models.INTERNAL_SERVER_ERROR
	mockUsecase.EXPECT().GetAll().Return(nil, err)

	var server *grpc.Server
	ctx := context.Background()
	req := &proto_grpc.EmptyRequest{}

	s := handler.NewPortServerGrpc(server, mockUsecase)
	p, err := s.GetAll(ctx, req)

	if err != nil || len(p.Ports) != 0 {
		t.Error("item should not be fetched in empty repository")
	}
}

func TestGetAllWithItems(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mock_usecase.NewMockPortUsecase(ctrl)

	port1 := &models.Port{
		ID: "ABCDE",
	}
	port2 := &models.Port{
		ID: "FGHIJ",
	}
	ports := []*models.Port{}
	ports = append(ports, port1)
	ports = append(ports, port2)

	mockUsecase.EXPECT().GetAll().Return(ports, nil)

	var server *grpc.Server
	ctx := context.Background()
	req := &proto_grpc.EmptyRequest{}

	s := handler.NewPortServerGrpc(server, mockUsecase)
	p, err := s.GetAll(ctx, req)

	if err != nil || len(p.Ports) == 0 {
		t.Error("item should be fetched in filled repository")
	}
}

func TestStoreNoItems(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mock_usecase.NewMockPortUsecase(ctrl)

	portRes := &proto_grpc.Port{
		ID: "ABCDE",
	}
	portReq := &models.Port{
		ID: "ABCDE",
	}
	mockUsecase.EXPECT().Store(portReq).Return(portReq, nil)

	var server *grpc.Server
	ctx := context.Background()

	s := handler.NewPortServerGrpc(server, mockUsecase)
	p, err := s.Store(ctx, portRes)

	if err != nil || p.ID != portReq.ID {
		t.Error("item should stored to empty repository")
	}
}

func TestStoreWithItems(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mock_usecase.NewMockPortUsecase(ctrl)

	portRes := &proto_grpc.Port{
		ID: "ABCDE",
	}
	portReq := &models.Port{
		ID: "ABCDE",
	}

	err := models.ITEM_EXIST_ERROR
	mockUsecase.EXPECT().Store(portReq).Return(nil, err)

	var server *grpc.Server
	ctx := context.Background()

	s := handler.NewPortServerGrpc(server, mockUsecase)
	p, err := s.Store(ctx, portRes)

	if err != models.ITEM_EXIST_ERROR || p != nil {
		t.Error("item should not be saved to repository, because item exists")
	}
}

func TestUpdateNoItems(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mock_usecase.NewMockPortUsecase(ctrl)

	portReq := &models.Port{
		ID: "ABCDE",
	}
	portRes := &proto_grpc.Port{
		ID: "ABCDE",
	}
	mockUsecase.EXPECT().Update(portReq).Return(portReq, nil)

	var server *grpc.Server
	ctx := context.Background()

	s := handler.NewPortServerGrpc(server, mockUsecase)
	p, err := s.Update(ctx, portRes)

	if err != nil || p.ID != portRes.ID {
		t.Error("existing item should be updated")
	}
}

func TestUpdateWithItems(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mock_usecase.NewMockPortUsecase(ctrl)

	portReq := &models.Port{
		ID: "ABCDE",
	}
	portRes := &proto_grpc.Port{
		ID: "ABCDE",
	}
	err := models.ITEM_EXIST_ERROR
	mockUsecase.EXPECT().Store(portReq).Return(portReq, nil)
	mockUsecase.EXPECT().Update(portReq).Return(nil, err)

	var server *grpc.Server
	ctx := context.Background()

	s := handler.NewPortServerGrpc(server, mockUsecase)

	_, storeErr := s.Store(ctx, portRes)
	if storeErr != nil {
		t.Error("failed saving item before testing update")
	}

	p, updateErr := s.Update(ctx, portRes)

	if updateErr != models.ITEM_EXIST_ERROR || p != nil {
		t.Error("item should not be updated, because item do not exists")
	}
}

func TestDeleteWithItems(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mock_usecase.NewMockPortUsecase(ctrl)

	portReq := &models.Port{
		ID: "ABCDE",
	}
	portRes := &proto_grpc.Port{
		ID: "ABCDE",
	}
	req := &proto_grpc.SingleRequest{
		Id: "ABCDE",
	}

	mockUsecase.EXPECT().Store(portReq).Return(portReq, nil)
	mockUsecase.EXPECT().Delete(req.Id).Return(true, nil)

	var server *grpc.Server
	ctx := context.Background()

	s := handler.NewPortServerGrpc(server, mockUsecase)

	_, storeErr := s.Store(ctx, portRes)
	if storeErr != nil {
		t.Error("failed saving item before testing update")
	}

	p, deleteErr := s.Delete(ctx, req)

	if deleteErr != nil || p.Code != 0 {
		t.Error("item should be deleted, because item was stored")
	}
}

func TestDeleteWithNoItems(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mock_usecase.NewMockPortUsecase(ctrl)

	req := &proto_grpc.SingleRequest{
		Id: "ABCDE",
	}

	err := models.NOT_FOUND_ERROR
	mockUsecase.EXPECT().Delete(req.Id).Return(false, err)

	var server *grpc.Server
	ctx := context.Background()

	s := handler.NewPortServerGrpc(server, mockUsecase)

	p, deleteErr := s.Delete(ctx, req)

	if deleteErr == nil || p.Code != -1000 {
		t.Error("item should not be deleted, because item was not stored")
	}
}

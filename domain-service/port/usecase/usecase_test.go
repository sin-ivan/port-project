package usecase_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	models "github.com/sin-ivan/port-project/domain-service/port"
	"github.com/sin-ivan/port-project/domain-service/port/repository/mocks"
	usecase "github.com/sin-ivan/port-project/domain-service/port/usecase"
)

//func SetupTest(

func TestGetByIDFromEmptyRepo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockPortRepository(ctrl)

	var err error = models.NOT_FOUND_ERROR
	mockRepo.EXPECT().GetByID("aaaaa").Return(nil, err)

	ucase := usecase.NewPortUsecase(mockRepo)
	p, err := ucase.GetByID("aaaaa")
	if err != models.NOT_FOUND_ERROR || p != nil {
		t.Error("port should not be found in empty repository")
	}
}

func TestGetByIDFromFromRepoWithItem(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockPortRepository(ctrl)

	port := &models.Port{
		ID: "ABCDE",
	}
	var err error = models.NOT_FOUND_ERROR
	mockRepo.EXPECT().GetByID("ABCDE").Return(port, nil)

	ucase := usecase.NewPortUsecase(mockRepo)
	p, err := ucase.GetByID("ABCDE")
	if err != nil || p == nil {
		t.Error("port should not be empty in repository")
	}
}

func TestGetAllFromEmptyRepo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockPortRepository(ctrl)

	var ports []*models.Port
	mockRepo.EXPECT().GetAll().Return(ports, nil)

	ucase := usecase.NewPortUsecase(mockRepo)
	p, err := ucase.GetAll()
	if err != nil || len(p) != 0 {
		t.Error("number of ports should not larger than 0")
	}
}

func TestGetAllFromRepoWithItems(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockPortRepository(ctrl)

	port1 := &models.Port{
		ID: "ABCDE",
	}
	port2 := &models.Port{
		ID: "FGHIJ",
	}

	ports := []*models.Port{}
	ports = append(ports, port1)
	ports = append(ports, port2)

	mockRepo.EXPECT().GetAll().Return(ports, nil)

	ucase := usecase.NewPortUsecase(mockRepo)
	p, err := ucase.GetAll()
	if err != nil || len(p) != 2 {
		t.Error("number of ports is not equal to provided values")
	}
}

func TestStoreToFilledInRepo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockPortRepository(ctrl)

	port := &models.Port{
		ID: "ABCDE",
	}

	err := models.ITEM_EXIST_ERROR
	mockRepo.EXPECT().GetByID("ABCDE").Return(port, nil)

	ucase := usecase.NewPortUsecase(mockRepo)
	p, err := ucase.Store(port)
	if err != models.ITEM_EXIST_ERROR || p != nil {
		t.Error("item should not be stored if it exists on repository")
	}
}

func TestStoreToEmptyRepo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockPortRepository(ctrl)

	port := &models.Port{
		ID: "ABCDE",
	}
	err := models.NOT_FOUND_ERROR
	mockRepo.EXPECT().GetByID("ABCDE").Return(nil, err)
	mockRepo.EXPECT().Store(port).Return(port.ID, nil)

	ucase := usecase.NewPortUsecase(mockRepo)
	p, err := ucase.Store(port)
	if err != nil || p.ID != port.ID {
		t.Error("item should be stored to empty repository")
	}
}

func TestStoreFailedRepo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockPortRepository(ctrl)

	port := &models.Port{
		ID: "ABCDE",
	}
	err := models.ITEM_EXIST_ERROR
	mockRepo.EXPECT().GetByID("ABCDE").Return(nil, err)
	mockRepo.EXPECT().Store(port).Return(port.ID, err)

	ucase := usecase.NewPortUsecase(mockRepo)
	p, err := ucase.Store(port)
	if err != models.ITEM_EXIST_ERROR || p != nil {
		t.Error("item should be stored to empty repository")
	}
}

func TestUpdateInEmptyRepo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockPortRepository(ctrl)

	port := &models.Port{
		ID: "ABCDE",
	}
	err := models.NOT_FOUND_ERROR
	mockRepo.EXPECT().GetByID("ABCDE").Return(nil, err)

	ucase := usecase.NewPortUsecase(mockRepo)
	p, err := ucase.Update(port)
	if err != models.NOT_FOUND_ERROR || p != nil {
		t.Error("item should not be updated in empty repository")
	}
}

func TestUpdateInFilledRepo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockPortRepository(ctrl)

	port := &models.Port{
		ID: "ABCDE",
	}
	mockRepo.EXPECT().GetByID("ABCDE").Return(port, nil)
	mockRepo.EXPECT().Update(port).Return(port, nil)

	ucase := usecase.NewPortUsecase(mockRepo)
	p, err := ucase.Update(port)
	if err != nil || p.ID != port.ID {
		t.Error("item should be updated in filled repository")
	}
}

func TestDeleteInEmptyRepo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockPortRepository(ctrl)

	port := &models.Port{
		ID: "ABCDE",
	}
	err := models.NOT_FOUND_ERROR
	mockRepo.EXPECT().GetByID("ABCDE").Return(nil, err)

	ucase := usecase.NewPortUsecase(mockRepo)
	r, err := ucase.Delete(port.ID)
	if err == nil || r {
		t.Error("item should not be remmoved from empty repository")
	}
}

func TestDeleteInFilledRepo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockPortRepository(ctrl)

	port := &models.Port{
		ID: "ABCDE",
	}
	mockRepo.EXPECT().GetByID("ABCDE").Return(port, nil)
	mockRepo.EXPECT().Delete(port.ID).Return(true, nil)

	ucase := usecase.NewPortUsecase(mockRepo)
	r, err := ucase.Delete(port.ID)
	if err != nil || !r {
		t.Error("item should removed from filled repository")
	}
}

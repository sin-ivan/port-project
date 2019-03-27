package repository_test

import (
	"fmt"
	"testing"

	models "github.com/sin-ivan/port-project/domain-service/port"
	repository "github.com/sin-ivan/port-project/domain-service/port/repository"
)

func TestStoreNoItems(t *testing.T) {
	repo := repository.NewMapPortRepository()

	port := &models.Port{
		ID: "ABCDE",
	}
	id, err := repo.Store(port)
	if id != port.ID || err != nil {
		t.Error("item should be store to empty map")
	}
}

func TestStoreWithItems(t *testing.T) {
	repo := repository.NewMapPortRepository()

	port := &models.Port{
		ID: "ABCDE",
	}
	_, err := repo.Store(port)
	if err != nil {
		t.Error("item should be stored to empty map")
	}
	id, err := repo.Store(port)
	if id != port.ID || err != models.ITEM_EXIST_ERROR {
		t.Error("item should be store to empty map")
	}
}

func TestGetByIDNoItems(t *testing.T) {
	repo := repository.NewMapPortRepository()

	port := &models.Port{
		ID: "ABCDE",
	}
	p, err := repo.GetByID(port.ID)
	if p != nil || err != models.NOT_FOUND_ERROR {
		t.Error("item should not be in empty map")
	}
}

func TestGetByIDWithItems(t *testing.T) {
	repo := repository.NewMapPortRepository()

	port := &models.Port{
		ID: "ABCDE",
	}
	_, err := repo.Store(port)
	if err != nil {
		t.Error("item should be stored to empty map")
	}

	p, err := repo.GetByID(port.ID)
	if p == nil || err != nil {
		t.Error("item should be found in map")
	}
}

func TestGetAllWithItems(t *testing.T) {
	repo := repository.NewMapPortRepository()

	for i := range [5]int{} {
		port := &models.Port{
			ID: fmt.Sprintf("ABCD%d", i),
		}
		_, err := repo.Store(port)
		if err != nil {
			t.Error("item should be stored to empty map")
		}
	}

	items, err := repo.GetAll()
	if len(items) != 5 || err != nil {
		t.Error("all items should be fetched")
	}
}

func TestGetAllNoItems(t *testing.T) {
	repo := repository.NewMapPortRepository()

	items, err := repo.GetAll()
	if len(items) != 0 || err != nil {
		t.Error("item list should be empty")
	}
}

func TestUpdateNoItems(t *testing.T) {
	repo := repository.NewMapPortRepository()

	port := &models.Port{
		ID: "ABCDE",
	}
	p, err := repo.Update(port)
	if p != nil || err != models.NOT_FOUND_ERROR {
		t.Error("item shoould not be updated in empty list")
	}
}

func TestUpdateWithItems(t *testing.T) {
	repo := repository.NewMapPortRepository()

	port := &models.Port{
		ID: "ABCDE",
	}
	_, err := repo.Store(port)
	if err != nil {
		t.Error("item should be stored to empty map")
	}

	port.City = "San-Francisco"
	p, err := repo.Update(port)
	if p == nil || p.City != "San-Francisco" || err != nil {
		t.Error("item shoould not be updated in empty list")
	}
}

func TestDeleteNoItems(t *testing.T) {
	repo := repository.NewMapPortRepository()

	port := &models.Port{
		ID: "ABCDE",
	}
	p, err := repo.Delete(port.ID)
	if p || err != models.NOT_FOUND_ERROR {
		t.Error("item shoould not be updated in empty list")
	}
}

func TestDeleteWithItems(t *testing.T) {
	repo := repository.NewMapPortRepository()

	port := &models.Port{
		ID: "ABCDE",
	}
	_, err := repo.Store(port)
	if err != nil {
		t.Error("item should be stored to empty map")
	}

	p, err := repo.Delete(port.ID)
	if !p || err != nil {
		t.Error("item shoould not be updated in empty list")
	}
}

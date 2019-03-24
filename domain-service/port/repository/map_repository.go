package repository

import (
	"sync"

	models "github.com/port-project/domain-service/port"
)

type mapPortRepository struct {
	storage sync.Map
}

// NewMapPortRepository is used to create repository instance
func NewMapPortRepository() PortRepository {
	// map is used for testing purposes
	return &mapPortRepository{sync.Map{}}
}

func (m *mapPortRepository) GetByID(id string) (*models.Port, error) {
	item, exist := m.storage.Load(id)
	if !exist {
		return nil, models.NOT_FOUND_ERROR
	}

	return item.(*models.Port), nil
}

func (m *mapPortRepository) GetAll() ([]*models.Port, error) {
	ports := make([]*models.Port, 0)

	m.storage.Range(func(k, v interface{}) bool {
		port := v.(*models.Port)
		ports = append(ports, port)
		return true
	})

	// error can be used later, with real DB instead of map
	return ports, nil
}

func (m *mapPortRepository) Store(p *models.Port) (string, error) {
	_, loaded := m.storage.Load(p.ID)
	if loaded {
		return p.ID, models.ITEM_EXIST_ERROR
	}

	// Store item if it was not found in DB
	m.storage.Store(p.ID, p)

	return p.ID, nil
}

func (m *mapPortRepository) Update(p *models.Port) (*models.Port, error) {
	_, loaded := m.storage.Load(p.ID)
	if !loaded {
		return nil, models.NOT_FOUND_ERROR
	}

	// Update item if it was found in storage
	m.storage.Store(p.ID, p)

	return p, nil
}

func (m *mapPortRepository) Delete(id string) (bool, error) {
	_, loaded := m.storage.Load(id)
	if !loaded {
		return false, models.NOT_FOUND_ERROR
	}

	// Remove item from storage
	m.storage.Delete(id)

	return true, nil
}

package repository

import models "github.com/sin-ivan/port-project/domain-service/port"

// PortRepository is used for persisting port data
type PortRepository interface {
	GetByID(string) (*models.Port, error)
	GetAll() ([]*models.Port, error)
	Store(p *models.Port) (string, error)
	Update(p *models.Port) (*models.Port, error)
	Delete(id string) (bool, error)
}

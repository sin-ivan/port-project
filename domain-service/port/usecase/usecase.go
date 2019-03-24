package usecase

import (
	models "github.com/port-project/domain-service/port"
	"github.com/port-project/domain-service/port/repository"
)

// PortUsecase is used for describing logic usecase
type PortUsecase interface {
	GetByID(id string) (*models.Port, error)
	GetAll() ([]*models.Port, error)
	Store(*models.Port) (*models.Port, error)
	Update(*models.Port) (*models.Port, error)
	Delete(id string) (bool, error)
}

type portUsecase struct {
	portRepos repository.PortRepository
}

// NewPortUsecase is used to create usecase instance
func NewPortUsecase(p repository.PortRepository) PortUsecase {
	return &portUsecase{p}
}

func (p *portUsecase) GetByID(id string) (*models.Port, error) {
	return p.portRepos.GetByID(id)
}

func (p *portUsecase) GetAll() ([]*models.Port, error) {
	return p.portRepos.GetAll()
}

func (p *portUsecase) Store(m *models.Port) (*models.Port, error) {
	port, _ := p.GetByID(m.ID)
	if port != nil {
		return nil, models.ITEM_EXIST_ERROR
	}

	_, err := p.portRepos.Store(m)
	if err != nil {
		return nil, err
	}

	// with real DB identifier should be generated after
	// data storage and then assigned to the Port object
	// like m.ID = id, in current case omitting this value

	return m, nil
}

func (p *portUsecase) Update(m *models.Port) (*models.Port, error) {
	port, err := p.GetByID(m.ID)
	if port != nil {
		return nil, err
	}

	return p.portRepos.Update(m)
}

func (p *portUsecase) Delete(id string) (bool, error) {
	article, _ := p.GetByID(id)
	if article == nil {
		return false, models.NOT_FOUND_ERROR
	}

	return p.portRepos.Delete(id)
}

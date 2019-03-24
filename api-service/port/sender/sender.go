package sender

import (
	models "github.com/port-project/api-service/port"
)

// PortSender is used for data submitting
type PortSender interface {
	StorePort(port *models.Port)
	GetAll() []*models.Port
}

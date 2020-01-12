package service

import (
	"gitlab.com/username/online-service-and-customer-care2.0/entity"
)

// ServiceRepository specifies companies's service related database operations
type ServiceRepository interface {
	Services() ([]entity.Service, []error)
	Service(name string) (*entity.Service, []error)
	//UpdateServie(service *entity.Service) (*entity.Service, []error)
	DeleteService(id uint) (*entity.Service, []error)
	StoreService(service *entity.Service) (*entity.Service, []error)
}
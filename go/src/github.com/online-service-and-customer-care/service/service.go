package service

import "gitlab.com/username/online-service-and-customer-care2.0/entity"

// ServiceService specifies company service related service
type ServiceService interface {
	Services() ([]entity.Service, []error)
	Service(name string) (*entity.Service, []error)
	//UpdateService(service *entity.Service) (*entity.Service, []error)
	DeleteService(id uint) (*entity.Service, []error)
	StoreService(service *entity.Service) (*entity.Service, []error)
}

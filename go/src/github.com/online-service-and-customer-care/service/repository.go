package Service


import (
	"onlineCustomerCare/entity"
)

// ServiceRepository specifies companies's service related database operations
type ServiceRepository interface {
	Services() ([]entity.Service, []error)
	Service(name string) (*entity.Service, []error)
	ServiceById(id int)(*entity.Service,[]error)
	UpdateService(service *entity.Service) (*entity.Service, []error)
	DeleteService(name string) (*entity.Service, []error)
	StoreService(service *entity.Service) (*entity.Service, []error)
}
package service_service

import (
	"gitlab.com/username/online-service-and-customer-care2.0/service"
	"gitlab.com/username/online-service-and-customer-care2.0/entity"
)

// ServiceService implements service.ServiceService interface
type ServiceService struct {
	servRepo service.ServiceRepository
}

// NewServiceService returns a new ServiceService object
func NewServiceService(servRepo service.ServiceRepository) *ServiceService {
	return &ServiceService{servRepo: servRepo}
}

// Services returns all stored services of the company
func (ss *ServiceService) Services() ([]entity.Service, []error) {
	srv, errs := ss.servRepo.Services()
	if len(errs) > 0 {
		return nil, errs
	}
	return srv, errs
}

// service retrieves stored service by its id
func (ss *ServiceService) Service(Name string) (*entity.Service, []error) {
	srv, errs := ss.servRepo.Service(Name)
	if len(errs) > 0 {
		return nil, errs
	}
	return srv, errs
}
/*
// UpdateService updates a given service
func (ss *ServiceService) UpdateService(service *entity.Service) (*entity.Service, []error) {
	srv, errs := ss.servRepo.UpdateService(service)
	if len(errs) > 0 {
		return nil, errs
	}
	return srv, errs
}
*/
// DeleteService deletes a given service
func (ss *ServiceService) DeleteServie(id uint) (*entity.Service, []error) {
	srv, errs := ss.servRepo.DeleteService(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return srv, errs
}

// StoreService stores a given service
func (ss *ServiceService) StoreService(service *entity.Service) (*entity.Service, []error) {
	srv, errs := ss.servRepo.StoreService(service)
	if len(errs) > 0 {
		return nil, errs
	}
	return srv, errs
}


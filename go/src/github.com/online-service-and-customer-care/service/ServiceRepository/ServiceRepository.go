package ServiceRepository

import (
	"github.com/jinzhu/gorm"
	"onlineCustomerCare/Service"
	"onlineCustomerCare/entity"
)

// ServiceGormRepo implements service.ServiceRepository interface
type ServiceGormRepo struct {
	conn *gorm.DB
}

// NewServiceGormRepo returns new object of ServiceGormRepo
func NewServiceGormRepo(db *gorm.DB) Service.ServiceRepository {
	return &ServiceGormRepo{conn: db}
}

// Services returns all company services stored in the database
func (servRepo *ServiceGormRepo) Services() ([]entity.Service, []error) {
	srv := []entity.Service{}
	errs := servRepo.conn.Find(&srv).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return srv, errs
}

// service retrieves a company's service from the database by its id
func (servRepo *ServiceGormRepo) Service(name string) (*entity.Service, []error) {
	srv := entity.Service{}
	errs := servRepo.conn.First(&srv, name).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &srv, errs
}

func (servRepo *ServiceGormRepo) ServiceById(id int) (*entity.Service, []error) {
	srv := entity.Service{}
	errs := servRepo.conn.Where("company_id = ?", id).First(&srv).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &srv, errs
}

// UpdateService updates a given company service in the database
func (servRepo *ServiceGormRepo) UpdateService(service *entity.Service) (*entity.Service, []error) {
	serv := service
	errs := servRepo.conn.Save(serv).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return serv, errs
}

// DeleteService deletes a given company's service  from the database
func (servRepo *ServiceGormRepo) DeleteService(name string) (*entity.Service, []error) {
	srv, errs := servRepo.Service(name)

	if len(errs) > 0 {
		return nil, errs
	}

	errs = servRepo.conn.Delete(srv, name).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return srv, errs
}

// StoreService stores a given company's service in the database
func (servRepo *ServiceGormRepo) StoreService(service *entity.Service) (*entity.Service, []error) {
	srv := service
	errs := servRepo.conn.Create(srv).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return srv, errs
}


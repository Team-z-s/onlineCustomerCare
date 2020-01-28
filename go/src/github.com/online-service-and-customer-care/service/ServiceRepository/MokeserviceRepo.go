package ServiceRepository

import (
	"errors"
	"onlineCustomerCare/Service"
	"onlineCustomerCare/entity"

	"github.com/jinzhu/gorm"
)

// ServiceGormRepo implements service.ServiceRepository interface
type MockServiceGormRepo struct {
	conn *gorm.DB
}

// NewServiceGormRepo returns new object of ServiceGormRepo
func NewMockServiceGormRepo(db *gorm.DB) Service.ServiceRepository {
	return &MockServiceGormRepo{conn: db}
}

// Services returns all company services stored in the database
func (mservRepo *MockServiceGormRepo) Services() ([]entity.Service, []error) {
	ctgs := []entity.Service{entity.MockService}
	return ctgs, nil
}

// service retrieves a company's service from the database by its id
func (mservRepo *MockServiceGormRepo) Service(name string) (*entity.Service, []error) {
	srv := entity.MockService
	if name == "mock service name" {
		return &srv, nil
	}
	return nil, []error{errors.New("Not found")}
}

func (mservRepo *MockServiceGormRepo) ServiceById(id int) (*entity.Service, []error) {
	srv := entity.MockService
	if id == 1 {
		return &srv, nil
	}
	return nil, []error{errors.New("not found ")}
}

// UpdateService updates a given company service in the database
func (mservRepo *MockServiceGormRepo) UpdateService(service *entity.Service) (*entity.Service, []error) {

	serv := entity.MockService
	return &serv, nil
}

// DeleteService deletes a given company's service  from the database
func (mservRepo *MockServiceGormRepo) DeleteService(name string) (*entity.Service, []error) {
	serv := entity.MockService
	if name != "mock service name" {
		return nil, []error{errors.New("Not found")}
	}
	return &serv, nil
}

// StoreService stores a given company's service in the database
func (mservRepo *MockServiceGormRepo) StoreService(service *entity.Service) (*entity.Service, []error) {
	srv := service

	return srv, nil
}

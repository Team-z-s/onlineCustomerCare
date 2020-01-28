package companyRepository

import (
	"errors"
	"onlineCustomerCare/company"
	"onlineCustomerCare/entity"

	"github.com/jinzhu/gorm"
)

type MockCompanyRepository struct {
	conn *gorm.DB
}

func NewMockCompanyRepository(c *gorm.DB) company.CompanyRepository {
	return &MockCompanyRepository{conn: c}
}

func (mcr *MockCompanyRepository) Companies() ([]entity.Companie, []error) {
	comp := []entity.Companie{entity.Companie{}}
	return comp ,nil
}
func (mcr *MockCompanyRepository) Company(id uint) (*entity.Companie, []error) {
	comp := entity.MockCompanie
	if id == 1 {
		return &comp, nil
	}
	return nil, []error{errors.New("Not found")}

}
func (mcr *MockCompanyRepository) CompanyByName(name string) (*entity.Companie, []error) {
	comp := entity.MockCompanie
	if name == "mock name" {
		return &comp, nil
	}
	return nil, []error{errors.New("Not found")}
}
func (mcr *MockCompanyRepository) UpdateCompany(company *entity.Companie) (*entity.Companie, []error) {
	comp := entity.MockCompanie
	return &comp, nil
}
func (mcr *MockCompanyRepository) DeleteCompany(id uint) (*entity.Companie, []error) {
	comp := entity.MockCompanie
	if id != 1 {
		return nil, []error{errors.New("Not found")}
	}
	return &comp, nil
}
func (mcr *MockCompanyRepository) StoreCompany(company *entity.Companie) (*entity.Companie, []error) {
	comp := company
	return comp, nil
}

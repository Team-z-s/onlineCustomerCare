package company

import "onlineCustomerCare/entity"

// CompanyService specifies company related service
type CompanyService interface {
	Companies() ([]entity.Companie, []error)
	Company(id uint) (*entity.Companie, []error)
	CompanyByName(name string) (*entity.Companie, []error)
	UpdateCompany(company *entity.Companie) (*entity.Companie, []error)
	DeleteCompany(id uint) (*entity.Companie, []error)
	StoreCompany(company *entity.Companie) (*entity.Companie, []error)
}

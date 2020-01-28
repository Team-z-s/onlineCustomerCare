package companyService

import ("onlineCustomerCare/company"
		"onlineCustomerCare/entity"
)


// CompanyService implements menu.CommentService interface
type CompanyService struct {
	companyRepo company.CompanyRepository
}

// NewCompanyService returns a new CompanyService object
func NewCompanyService(compRepo company.CompanyRepository) company.CompanyService {
	return &CompanyService{companyRepo: compRepo}
}

// Companies returns all stored companies
func (cs *CompanyService) Companies() ([]entity.Companie, []error) {
	comps, errs := cs.companyRepo.Companies()
	if len(errs) > 0 {
		return nil, errs
	}
	return comps, errs
}

// Company retrieves stored company by its id
func (cs *CompanyService) Company(id uint) (*entity.Companie, []error) {
	comp, errs := cs.companyRepo.Company(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return comp, errs
}
// Company retrieves stored company by its id
func (cs *CompanyService) CompanyByName(name string) (*entity.Companie, []error) {
	comp, errs := cs.companyRepo.CompanyByName(name)
	if len(errs) > 0 {
		return nil, errs
	}
	return comp, errs
}
// UpdateCompany updates a given company
func (cs *CompanyService) UpdateCompany(company *entity.Companie) (*entity.Companie, []error) {
	comp, errs := cs.companyRepo.UpdateCompany(company)
	if len(errs) > 0 {
		return nil, errs
	}
	return comp, errs
}

// DeleteCompany deletes a given comment
func (cs *CompanyService) DeleteCompany(id uint) (*entity.Companie, []error) {
	comp, errs := cs.companyRepo.DeleteCompany(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return comp, errs
}

// StoreComment stores a given comment
func (cs *CompanyService) StoreCompany(company *entity.Companie) (*entity.Companie, []error) {
	comp, errs := cs.companyRepo.StoreCompany(company)
	if len(errs) > 0 {
		return nil, errs
	}
	return comp, errs
}



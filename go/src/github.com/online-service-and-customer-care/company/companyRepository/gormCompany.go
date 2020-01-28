package companyRepository

import ("onlineCustomerCare/company"
		"onlineCustomerCare/entity"
	"github.com/jinzhu/gorm"
)


// CompanyGormRepo implements company repository interface
type CompanyGormRepo struct {
	conn *gorm.DB
}

// NewCompanyGormRepo returns new object of CompanyGormRepo
func NewCompanyGormRepo(db *gorm.DB) company.CompanyRepository {
	return &CompanyGormRepo{conn: db}
}

// Companies returns all companies stored in the database
func (compRepo *CompanyGormRepo) Companies() ([]entity.Companie, []error) {
	comps := []entity.Companie{}
	//fmt.Println(cmnts)
	errs := compRepo.conn.Find(&comps).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return comps, errs
}

// Company retrieves a company from the database by its id
func (compRepo *CompanyGormRepo) Company(id uint) (*entity.Companie, []error) {
	comps := entity.Companie{}
	errs := compRepo.conn.First(&comps, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &comps, errs
}
// CompanyByName retrieves a company from the database by its name
func (compRepo *CompanyGormRepo) CompanyByName(name string) (*entity.Companie, []error) {
	comps := entity.Companie{}
	errs := compRepo.conn.Where("full_name=?",name).First(&comps).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &comps, errs
}

// UpdateCompany updates a given company in the database
func (compRepo *CompanyGormRepo) UpdateCompany(company *entity.Companie) (*entity.Companie, []error) {
	comp := company
	errs := compRepo.conn.Save(comp).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return comp, errs
}

// DeleteCompany deletes a given company from the database
func (compRepo *CompanyGormRepo) DeleteCompany(id uint) (*entity.Companie, []error) {
	comp, errs := compRepo.Company(id)

	if len(errs) > 0 {
		return nil, errs
	}

	errs = compRepo.conn.Delete(comp, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return comp, errs
}

// StoreCompany stores a given company in the database
func (compRepo *CompanyGormRepo) StoreCompany(company *entity.Companie) (*entity.Companie, []error) {
	comp := company
	errs := compRepo.conn.Save(comp).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return comp, errs
}


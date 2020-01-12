package login_repository


import (
	"github.com/jinzhu/gorm"
	"gitlab.com/username/online-service-and-customer-care2.0/entity"
	"gitlab.com/username/online-service-and-customer-care2.0/login"
)

//LoginGormRepo implments service of Login
type LoginGormRepo struct {
	conn *gorm.DB
}
//NewLoginGormRepo creates instance of LoginGromRepo
func NewLoginGormRepo(db *gorm.DB) login.LoginRepository{
	return &LoginGormRepo{conn: db}
}

// users returns all user stored in the database
func (logRepo *LoginGormRepo) Users() ([]entity.User, []error) {
	us := []entity.User{}
	errs := logRepo.conn.Find(&us).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return us, errs
}
// Employees returns all Employees stored in the database
func (logRepo *LoginGormRepo) Employees() ([]entity.Employee, []error) {
	empls := []entity.Employee{}
	errs := logRepo.conn.Find(&empls).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return empls, errs
}

// Companies returns all companies stored in the database
func (logRepo *LoginGormRepo) Companies() ([]entity.Companie, []error) {
	comps := []entity.Companie{}

	errs := logRepo.conn.Find(&comps).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return comps, errs
}

// Roles returns all Roles stored in the database
func (logRepo *LoginGormRepo) Roles() ([]entity.Roles, []error) {
	comps := []entity.Roles{}

	errs := logRepo.conn.Find(&comps).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return comps, errs
}
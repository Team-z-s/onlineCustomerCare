package login_service

import (
	
	"gitlab.com/username/online-service-and-customer-care2.0/entity"
	"gitlab.com/username/online-service-and-customer-care2.0/login"
)

//LoginserviceGorm implments service of Login
type LoginServiceGorm struct {
	logRepo login.LoginRepository
}

//NewEmployeeServiceGorm construstor
func NewLoginServiceGorm(logRepo login.LoginRepository) login.LoginService {
	return &LoginServiceGorm{logRepo: logRepo}
}
//Employees return all employees on the database
func (ls *LoginServiceGorm) Employees() ([]entity.Employee, []error) {
	empls, errs := ls.logRepo.Employees()
	if len(errs) > 0 {
		return nil, errs
	}
	return empls, errs
}
func (ls *LoginServiceGorm) Companies() ([]entity.Companie, []error) {
	comps, errs := ls.logRepo.Companies()
	if len(errs) > 0 {
		return nil, errs
	}
	return comps, errs
}

// Users returns all stored users
func (ls *LoginServiceGorm) Users() ([]entity.User, []error) {
	usrs, errs := ls.logRepo.Users()
	if len(errs) > 0 {
		return nil, errs
	}
	return usrs, errs
}

// Users returns all stored users
func (ls *LoginServiceGorm) Roles() ([]entity.Roles, []error) {
	usrs, errs := ls.logRepo.Roles()
	if len(errs) > 0 {
		return nil, errs
	}
	return usrs, errs
}
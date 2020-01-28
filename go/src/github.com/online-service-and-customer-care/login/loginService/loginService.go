package loginService

import (
	"onlineCustomerCare/login"
	"onlineCustomerCare/entity"
)

type LoginServiceGorm struct{
	LogRepo login.LoginRepository
}
//NewLoginServiceGorm construstor
func NewLoginServiceGorm(logRepo login.LoginRepository) login.LoginService {
	return &LoginServiceGorm{LogRepo: logRepo}
}

//Account return a single account
func (ls *LoginServiceGorm) Account(username string) (*entity.Account, error) {
	account, errs := ls.LogRepo.Account(username)
	if  errs != nil{
		return nil, errs
	}
	return account, errs
}

//UpdateAccount updates
func (ls *LoginServiceGorm) UpdateAccount(act *entity.Account) (*entity.Account, []error) {
	account, errs := ls.LogRepo.UpdateAccount(act)
	if len(errs) > 0 {
		return nil, errs
	}
	return account, errs
}

//DeleteAccount delete
func (ls *LoginServiceGorm) DeleteAccount(username string) (*entity.Account, []error) {
	account, errs := ls.LogRepo.DeleteAccount(username)
	if len(errs) > 0 {
		return nil, errs
	}
	return account, errs
}

//StoreAccount store
func (ls *LoginServiceGorm) StoreAccount(account *entity.Account) (*entity.Account, []error) {
	act, errs := ls.LogRepo.StoreAccount(account)
	if len(errs) > 0 {
		return nil, errs
	}
	return act, errs
}



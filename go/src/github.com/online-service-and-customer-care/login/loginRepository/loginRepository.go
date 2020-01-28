package loginRepository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"onlineCustomerCare/login"
	"onlineCustomerCare/entity"
)

//LoginGromRepo implments Login repository
type LoginGormRepo struct {
	dbconn *gorm.DB
}

//NewLoginGromRepo isa construtor for LoginGromRepo+
func NewLoginGormRepo(db *gorm.DB) login.LoginRepository {
	return &LoginGormRepo{dbconn: db}
}

// Account retrieves an account from the database by its username
func (logrepo *LoginGormRepo) Account(username string) (*entity.Account, error) {
	account := entity.Account{}
	errs := logrepo.dbconn.Where("username=?",username).First(&account).Error

	if errs != nil {
		fmt.Println("*************************************************************************8")
		return nil, errs
	}
	return &account, errs
}

// UpdateAccount updates a given Account in the database
func (logRepo *LoginGormRepo) UpdateAccount(account *entity.Account) (*entity.Account, []error) {
	act := account
	errs := logRepo.dbconn.Save(act).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return act, errs
}

// DeleteAccount deletes a given Account from the database
func (LogRepo *LoginGormRepo) DeleteAccount(username string) (*entity.Account, []error) {
	account, errs := LogRepo.DeleteAccount(username)

	if len(errs) > 0 {
		return nil, errs
	}

	errs = LogRepo.dbconn.Delete(account, username).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return account, errs
}

// StoreAccount stores an Account in the database
func (LogRepo *LoginGormRepo) StoreAccount(act *entity.Account) (*entity.Account, []error) {
	account := act
	errs := LogRepo.dbconn.Create(account).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return account, errs
}


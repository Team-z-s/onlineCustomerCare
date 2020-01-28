package loginRepository

import (
	"errors"
	"onlineCustomerCare/entity"
	"onlineCustomerCare/login"

	"github.com/jinzhu/gorm"
)

//LoginGromRepo implments Login repository
type MockLoginGormRepo struct {
	dbconn *gorm.DB
}

//NewLoginGromRepo isa construtor for LoginGromRepo+
func MockNewLoginGormRepo(db *gorm.DB) login.LoginRepository {
	return &MockLoginGormRepo{dbconn: db}
}

// Account retrieves an account from the database by its username
func (logrepo *MockLoginGormRepo) Account(username string) (*entity.Account, error) {
	account := entity.MockAccount
	if username == "mock account username"{
		return &account, nil
	}
	return nil, errors.New("not found")

}

// UpdateAccount updates a given Account in the database
func (logrepo *MockLoginGormRepo) UpdateAccount(account *entity.Account) (*entity.Account, []error) {
	act := entity.MockAccount
	return &act, nil
}

// DeleteAccount deletes a given Account from the database
func (logrepo *MockLoginGormRepo) DeleteAccount(username string) (*entity.Account, []error) {
	account := entity.MockAccount
	if username == ""{
		return &account,nil
	}
	return nil, []error{errors.New("not found")}
}

// StoreAccount stores an Account in the database
func (logrepo *MockLoginGormRepo) StoreAccount(act *entity.Account) (*entity.Account, []error) {
	account := entity.MockAccount
	
	return &account, nil
}

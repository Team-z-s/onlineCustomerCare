package login

import "onlineCustomerCare/entity"

// LoginService specifies application employee related services
type LoginService interface {
	Account(username string)(*entity.Account,error)
	UpdateAccount(account *entity.Account) (*entity.Account, []error)
	DeleteAccount(username string) (*entity.Account, []error)
	StoreAccount(account *entity.Account) (*entity.Account, []error)
}
type SessionService interface {
	Session(sessionID string) (*entity.Session, []error)
	StoreSession(session *entity.Session) (*entity.Session, []error)
	DeleteSession(sessionID string) (*entity.Session, []error)
}
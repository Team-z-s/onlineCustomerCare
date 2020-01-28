package login

import "onlineCustomerCare/entity"

// LoginService specifies login related services
type LoginRepository interface {
	Account(username string)(*entity.Account,error)
	UpdateAccount(account *entity.Account) (*entity.Account, []error)
	DeleteAccount(username string) (*entity.Account, []error)
	StoreAccount(account *entity.Account) (*entity.Account, []error)
}
type SessionRepository interface {
	Session(sessionID string) (*entity.Session, []error)
	StoreSession(session *entity.Session) (*entity.Session, []error)
	DeleteSession(sessionID string) (*entity.Session, []error)
}
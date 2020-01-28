package user

import "onlineCustomerCare/entity"

//UserRepository specifies db opration for user
type UserRepository interface {
	Users() ([]entity.User, []error)
	User(id uint) (*entity.User, []error)
	UpdateUser(user *entity.User) (*entity.User,[]error)
	DeleteUser(id uint) (*entity.User,[]error)
	StoreUser(user *entity.User) (*entity.User,[]error)
}


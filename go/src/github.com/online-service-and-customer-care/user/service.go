package user

import (
	"gitlab.com/username/online-service-and-customer-care2.0/entity"
)

// UserService contain all the service of user
type UserService interface {
	Users() ([]entity.User, []error)
	User(id uint) (*entity.User, []error)
	UpdateUser(user *entity.User)(*entity.User, []error)
	DeleteUser(id uint) (*entity.User, []error)
	StoreUser(user *entity.User) (*entity.User,[]error)
}

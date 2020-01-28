package userService

import (
	"onlineCustomerCare/entity"
	"onlineCustomerCare/user"
)

// UserService implements user.UserService interface
type UserService struct {
	userRepo user.UserRepository
}

// NewUserService returns a new instance of UserService
func NewUserService(userRepo user.UserRepository) user.UserService {
	return &UserService{userRepo: userRepo}
}

// Users returns all stored users
func (us *UserService) Users() ([]entity.User, []error) {
	usrs, errs := us.userRepo.Users()
	if len(errs) > 0 {
		return nil, errs
	}
	return usrs, errs
}

// User retrieves stored User by its id
func (us *UserService) User(id uint) (*entity.User, []error) {
	usr, errs := us.userRepo.User(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

// UpdateUser updates a given user in the database
func (us *UserService) UpdateUser(user *entity.User) (*entity.User, []error) {
	usr, errs := us.userRepo.UpdateUser(user)
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

// DeleteUser deletes a given user
func (us *UserService) DeleteUser(id uint) (*entity.User, []error) {
	user, errs := us.userRepo.DeleteUser(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return user, errs
}

// StoreUser stores a given user
func (us *UserService) StoreUser(user *entity.User) (*entity.User, []error) {
	return nil, nil
}


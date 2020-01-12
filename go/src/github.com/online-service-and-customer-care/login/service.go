package login

import "gitlab.com/username/online-service-and-customer-care2.0/entity"

// LoginService specifies application employee related services
type LoginService interface {
	Users() ([]entity.User, []error)
	Companies() ([]entity.Companie, []error)
	Employees() ([]entity.Employee, []error)
	Roles() ([]entity.Roles, []error)
}

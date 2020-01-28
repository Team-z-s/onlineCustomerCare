package employeeRepository
/*
import (
	"errors"
	"onlineCustomerCare/employee"
	"onlineCustomerCare/entity"

	"github.com/jinzhu/gorm"
)

type MockEmployeeGormRepo struct {
	conn *gorm.DB
}

func NewMockEmployeeGormRepo(db *gorm.DB) *employee.EmployeeRepository {
	return &MockEmployeeGormRepo{conn: db}
}

func (meRepo *MockEmployeeGormRepo) Employees() ([]entity.Employee, []error) {
	empl := []entity.Employee{entity.MockEmployee}
	return empl, nil
}

func (meRepo *MockEmployeeGormRepo) Employee(id uint) (*entity.Employee, []error) {
	empl := entity.MockEmployee
	if id == 1 {
		return &empl, nil
	}
	return nil, []error{errors.New("Not found")}
}

func (meRepo *MockEmployeeGormRepo) UpdateEmployee(user *entity.Employee) (*entity.Employee, []error) {

	empl := entity.MockEmployee
	return &empl, nil
}

func (meRepo *MockEmployeeGormRepo) DeleteEmployee(id uint) (*entity.Employee, []error) {
	empl := entity.MockEmployee
	if id != 1 {
		return nil, []error{errors.New("Not found")}
	}
	return &empl, nil
}

func (meRepo *MockEmployeeGormRepo) StoreEmployee(user *entity.Employee) (*entity.Employee, []error) {
	empl := user

	return empl, nil
}
*/
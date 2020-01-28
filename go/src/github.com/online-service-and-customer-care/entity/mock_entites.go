package entity

//Admin struct contain admin info
type mockAdmin struct {
	username string
	password string
}

//User struct contain info about user
var MockUser = User{
	ID:       1,
	FName:    "mock user first name",
	LName:    "mock user last name",
	Email:    "mock user email",
	Phone:    "mock phone number",
	Username: "mock username",
	Password: "mock password",
}

//Employee struct contains employees
var MockEmployee = Employee{
	EmployeeID: 1,
	CompanyID:  1,
	FName:      "mock employee first name",
	LName:      "mock employee last name",
	Email:      "mock employee email",
	Address:    "mock imployee adress",
	Salary:     20000.4,
	Phone:      "mock emoloyee phone",
	Photo:      "mock employee photo",
	Username:   "mock employee username",
	Password:   "mock employee password",
}

//Company struct contain fields of the company
var MockCompanie = Companie{
	CompanyID: 1,
	FullName:  "company name",
	Logo:      "logo",
	Email:     "company email",
	Phone:     "phone number",
	Address:   "address",
	Moto:      "moto",
	Password:  "password",
}

//Service contains fields of service
var MockService = Service{
	ID:          1,
	CompanyID:   1,
	Name:        "mock service name",
	Description: "mock description",
}

//Roles role
var MockRoles = Roles{
	ID:   2,
	Name: "company",
}

var MockAccount = Account{
	Username: "mock account username",
	Password: "mock account password",
	Role_id:  2,
}

//Task taskes assigned for the employee

type MockTask struct {
	ID          int
	EmployeeID  int
	Employee    Employee
	Name        string `gorm:"type:varchar(255);not null"`
	Description string
	Progress    string `gorm:"type:varchar(255);"`
}

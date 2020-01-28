package entity

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)



var JwtKey = []byte("my_secret_key")

type Claims struct {
	Username       string `json:"username"`
	jwt.StandardClaims
}

//Admin struct contain admin info
type Admin struct {
	username string
	password string
}

//User struct contain info about user
type User struct {

	ID int
	FName string `gorm:"type:varchar(255); not null"`
	LName string `gorm:"type:varchar(255); not null"`
	Email string `gorm:"type:varchar(255);unique"`
	Phone string `gorm:"type:varchar(255)"`

	Username string `gorm:"type:varchar(255); not null; unique"`
	Password string `gorm:"type:varchar(255); not null"`
}

//Address contains country,city, woreda,kebele
type Address struct {
	country, city, woreda, kebele string
}

//Employee struct contains employees
type Employee struct {
	EmployeeID        int
	CompanyID int
	company Companie
	FName     string  `gorm:"type:varchar(255)" `
	LName     string  `gorm:"type:varchar(255)"`
	Email     string  `gorm:"type:varchar(255);not null; unique"`
	Address   string
	Salary    float64
	Phone     string  `gorm:"type:varchar(255)"`
	Photo     string  `gorm:"type:varchar(255)"`
	Username  string  `gorm:"type:varchar(255);not null; unique"`
	Password  string  `gorm:"type:varchar(255); not null "`
}

//Company struct contain fields of the company
type Companie struct {
	CompanyID       int
	FullName string `gorm:"type:varchar(255);not null; unique"`
	Logo     string
	Email    string `gorm:"type:varchar(255);not null; unique"`
	Phone    string `gorm:"type:varchar(255);not null; unique"`
	Address  string
	Moto     string
	Password string `gorm:"type:varchar(255);not null "`
}

//Service contains fields of service
type Service struct {
	ID          int
	CompanyID   int
	Company Companie
	Name        string `gorm:"type:varchar(255);not null "`
	Description string
}

//Date c
type Date struct {
	Day   int
	Month int
	Year  int
}

//Attendance contain date and time
type Attendance struct {
	ID         int       `json:"id"`
	EmployeeID int       `json:"employee_id"`
	Shift      string    `json:"shift"`
	Date       Date      `json:"data_a"`
	Status     string    `json:"status"`
	PostedAt   time.Time `json:"posted_at"`
}

//Comment comment
type Comment struct {
	ID        uint
	FullName  string `gorm:"type:varchar(255)"`
	Message   string
	Phone     string `gorm:"type:varchar(100);not null; unique"`
	Email     string `gorm:"type:varchar(255);not null; unique"`
	CreatedAt time.Time
}
//Jobs employee's work
type Employee_job struct {
	EmployeeID int
	CompanyID  int
	Employee Employee
	Job        string  `json:"job"`
	DayRate    float64 `json:"day_rate"`
	NightRate  float64 `json:"night_rate"`
}

//Roles role
type Roles struct {
	ID   int
	Name string
}

type Account struct{
	Username string `gorm:"type:varchar(255);not null;unique"`
	Password string `gorm:"type:varchar(255);not null"`

	Role_id int `json:"role_id" `
}

//Task taskes assigned for the employee

type Task struct {
	ID          int
	EmployeeID  int
	Employee Employee
	Name        string `gorm:"type:varchar(255);not null"`
	Description string
	Progress    string `gorm:"type:varchar(255);"`
}
//Session represents login user session
type Session struct {
	//ID         uint
	//Uername    string `gorm:"type:varchar(255);not null"`
	//UUID       string `gorm:"type:varchar(255);not null"`
	//Expires    int64  `gorm:"type:varchar(255);not null"`
	SigningKey []byte `gorm:"type:varchar(255);not null"`
}

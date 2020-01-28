package main

import (
	"fmt"
	"onlineCustomerCare/entity"
	"onlineCustomerCare/login/loginRepository"
	"onlineCustomerCare/login/loginService"
	"onlineCustomerCare/rtoken"
	"onlineCustomerCare/session"
	"onlineCustomerCare/user/userRepository"
	"onlineCustomerCare/user/userService"
	//"time"

	//"github.com/gorilla/mux"
	//"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"net/http"
	"onlineCustomerCare/company/companyRepository"
	"onlineCustomerCare/company/companyService"
	"onlineCustomerCare/employee/employeeRepository"
	"onlineCustomerCare/employee/employeeService"
	"onlineCustomerCare/Service/ServiceRepository"
	"onlineCustomerCare/Service/ServiceService"
	"onlineCustomerCare/delivery/http/handler"
	"text/template"
)
var tmpl = template.Must(template.ParseGlob("C:/Users/hp/go/src/onlineCustomerCare/ui/templates/*"))
func Index(w http.ResponseWriter,r*http.Request ){
	tmpl.ExecuteTemplate(w,"index.html",nil)
}
func createTables(dbconn *gorm.DB) []error {
	errs := dbconn.CreateTable(&entity.User{}, &entity.Roles{}, &entity.Session{}, &entity.Employee{}, &entity.Employee_job{}, &entity.Companie{}, &entity.Service{},&entity.Account{}, &entity.Comment{}).GetErrors()
	if errs != nil {
		return errs
	}
	return nil
}

func main() {
	//createTables(dbconn)

	csrfSignKey := []byte(rtoken.GenerateRandomID(32))

	fmt.Println(csrfSignKey)

	tmpl := template.Must(template.ParseGlob("C:/Users/hp/go/src/onlineCustomerCare/ui/templates/*"))

	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", "postgres",
		"lealem-g","localhost", "customercare")
	dbconn, err := gorm.Open("postgres", connStr)



	if err != nil  {
		panic(err)
	}

	defer dbconn.Close()
//	commentRepo := commentRepository.NewCommentGormRepo(dbconn)
//	commentService := commentService.NewCommentService(commentRepo)
	//commentHand := handler.NewCommentHandler(commentService)



	logrepo := loginRepository.NewLoginGormRepo(dbconn)
	logserv := loginService.NewLoginServiceGorm(logrepo)
	//sessionrepo := loginRepository.NewSessionGormRepo(dbconn)
	//sessionsrv := loginService.NewSessionService(sessionrepo)

	loghand := handler.NewLoginHandler(logserv,tmpl)

	userRepo := userRepository.NewUserGormRepo(dbconn)
	userServ := userService.NewUserService(userRepo)
	userHand := handler.NewUserHandler(userServ,tmpl)

	compRepo := companyRepository.NewCompanyGormRepo(dbconn)
	compserv := companyService.NewCompanyService(compRepo)
	comphand := handler.NewCompanyHandler(compserv,logserv,tmpl)

	sercRepo := ServiceRepository.NewServiceGormRepo(dbconn)
	servServ := ServiceService.NewServiceService(sercRepo)
	serviceHand := handler.NewServiceHandler(servServ,compserv,tmpl)

	emplRepo := employeeRepository.NewEmployeeGormRepo(dbconn)
	emplSerc := employeeService.NewEmployeeServiceGorm(emplRepo)
	emplHand := handler.NewEmloyeeHandler(emplSerc,logserv,compserv,tmpl)

//##############################################mulriplexer and http handle functions######################################################################

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

//*************************************************************************************

	mux.HandleFunc("/",Index)  //+++++++++++++++++++++++
	mux.HandleFunc("/login",loghand.Login) //++++++++++++++++++++++
	mux.HandleFunc("/user/signUp",userHand.SignUp) //++++++++++++++++++++++
	mux.HandleFunc("/logout",loghand.Logout)
	mux.HandleFunc("/user/search",comphand.Search)
	mux.HandleFunc("/company/detail",comphand.GetCompany)
	//mux.HandleFunc("/user/postComment",commentHand.AddComment)

	mux.HandleFunc("/admin_dashboard",comphand.AdminDashboard) //+++++++++++++++++++++++++
	mux.HandleFunc("/company_dashboard",loghand.CompanyDashboard) //+++++++++++++++++++++++

	mux.Handle("/employee/profile",session.Authenticated(http.HandlerFunc(emplHand.ShowProfile)))  //++++++++++++--------------
	mux.Handle("/employee/seeTask",session.Authenticated(http.HandlerFunc(emplHand.AssignTask))) //---------


	mux.Handle("/company/addEmployee",session.Authenticated(http.HandlerFunc(emplHand.AddEmployee)))//++++++++++++++++++++++++--
	mux.Handle("/company/deleteEmployee",session.Authenticated(http.HandlerFunc(emplHand.DeleteEmployee)))//+++++++++-------------
	//mux.HandleFunc("/company/updateEmployee",emplHand.UpdateEmployee)
	mux.HandleFunc("/company/update",comphand.UpdateCompany)
	mux.HandleFunc("/company/assignTask",emplHand.AssignTask)
	mux.Handle("/company/addService",session.Authenticated(http.HandlerFunc(serviceHand.AddService)))//++++++++++++++++++++++---------
	mux.Handle("/company/deleteService",session.Authenticated(http.HandlerFunc(serviceHand.DeleteService)))
	mux.Handle("/company/updateService",session.Authenticated(http.HandlerFunc(serviceHand.UpdateService)))
	//mux.HandleFunc("/company/trackProgress",emplHand.Track)

	mux.Handle("/admin/addCompany",session.Authenticated(http.HandlerFunc(comphand.AddCompany)))//+++++++++++++++++++--
	mux.Handle("/admin/deleteCompany",session.Authenticated(http.HandlerFunc(comphand.DeleteCompany)))//++++++++++++++++--



	http.ListenAndServe(":8282",mux)


}




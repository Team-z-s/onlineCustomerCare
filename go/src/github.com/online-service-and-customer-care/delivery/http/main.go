package main

import (
	
	"net/http"
"html/template"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/julienschmidt/httprouter"

	"gitlab.com/username/online-service-and-customer-care2.0/comment/Crepository"
	"gitlab.com/username/online-service-and-customer-care2.0/comment/cservice"
	"gitlab.com/username/online-service-and-customer-care2.0/user/user_repository"
	"gitlab.com/username/online-service-and-customer-care2.0/user/user_service"
	"gitlab.com/username/online-service-and-customer-care2.0/login/login_repository"
	"gitlab.com/username/online-service-and-customer-care2.0/login/login_service"
	"gitlab.com/username/online-service-and-customer-care2.0/delivery/http/handler"
	"gitlab.com/username/online-service-and-customer-care2.0/employee/employee_repository"
	"gitlab.com/username/online-service-and-customer-care2.0/employee/employee_service"
	"gitlab.com/username/online-service-and-customer-care2.0/company/company_service"
	"gitlab.com/username/online-service-and-customer-care2.0/company/company_repository"
)




func main() {
	tmpl := template.Must(template.ParseGlob("../../ui/templates/*"))
	dbconn, err := gorm.Open("postgres", "postgres://postgres:lealem-g@localhost/customercare?sslmode=disable")

	if err != nil {
		panic(err)
	}

	defer dbconn.Close()

	//employee related repository and service
	emplRepo := employee_repository.NewEmployeeGormRepo(dbconn)
	emplSrv := employee_service.NewEmployeeServiceGorm(emplRepo)
	employeeHandler := handler.NewEmployeeHandler(emplSrv,tmpl)
	//comment related repository and service
	comRepo := Crepository.NewCommentGormRepo(dbconn)
	comserv := cservice.NewCommentService(comRepo)
	comhan := handler.NewAdminCommentHandler(comserv,tmpl)

	//login related repository and service
	logRepo := login_repository.NewLoginGormRepo(dbconn)
	logserv := login_service.NewLoginServiceGorm(logRepo)
	loginhandler := handler.NewLoginHandler(logserv,tmpl)

	//user related repository and service
	userRepo := user_repository.NewUserGormRepo(dbconn)
	userserv := user_service.NewUserService(userRepo)
	userhandler := handler.NewUserHandler(userserv,tmpl)

	//company related repository and service
	compRepo := company_repository.NewCompanyGormRepo(dbconn)
	compServ := company_service.NewCompanyService(compRepo)
	comphandler := handler.NewCompanyHandler(compServ,tmpl)

	//admin company  related repository and service
	admincompanyhandler := handler.NewAdminCompanyHandler(compServ,tmpl)

	router := httprouter.New()

	router.GET("/admin_dashboard",loginhandler.Admindashboard)
	router.POST("/admin_dashboard/add_Company",admincompanyhandler.AddCompany)
	router.DELETE(".admin_dashboard/delete_company",admincompanyhandler.DeleteCompany)

	router.GET("/",loginhandler.Index)
	router.POST("/login",loginhandler.Login)
	router.POST("/user/signup",userhandler.Signin)
	router.GET("/user/search",userhandler.Search)
	router.POST("/user/comment",comhan.PostComment)

	router.GET("/company/comment",comhan.GetComments)
	router.POST("/company_dashboard",loginhandler.Companydashboard)
	router.POST("company_dashboard/changelogo",comphandler.ChangeLogo)
	router.POST("company_dashboard/addservice",comphandler.Addservice)
	router.POST("/company_dashboard/take_attendance",comphandler.Take_attendance)
	router.POST("/company_dashboard/calculatesallary",comphandler.Calculate_salary)
	router.POST("/company_dashboard/add_employee", employeeHandler.PostEmployee)
	router.DELETE("/company_dashboard/add_employee:id", employeeHandler.DeleteEmployee)

	http.ListenAndServe(":8181", router)

}

package handler

import (
	"github.com/julienschmidt/httprouter"
	"gitlab.com/username/online-service-and-customer-care2.0/entity"
	"gitlab.com/username/online-service-and-customer-care2.0/login"
	"html/template"
	"net/http"
)
type LoginHandler struct {
	logService login.LoginService
	tmpl        *template.Template
}
// NewEmployeeHandler returns new AdminCommentHandler object
func NewLoginHandler(logService login.LoginService, T *template.Template) *LoginHandler {
	return &LoginHandler{logService: logService, tmpl:T}
}
//Login check if the user is autoraized
func (lh *LoginHandler) Login(w http.ResponseWriter,r *http.Request, _ httprouter.Params){
	if r.Method == http.MethodPost{
		var username = r.FormValue("username")
		var password = r.FormValue("password")

		if username == "teamz" && password =="teamzpass"{
		 	http.Redirect(w, r, "/admin_dashboard", http.StatusSeeOther)
		 }
 		var entry =lh.chakeRole(username,password)
 		if entry == "user"{

			http.Redirect(w,r,"/user/search",http.StatusSeeOther)

	//	http.Redirect(w, r, "/admin_dashboard", http.StatusSeeOther) 
 } else if entry == "company"{
	http.Redirect(w, r, "/company_dashboard", http.StatusSeeOther)
	//lh.tmpl.ExecuteTemplate(w,"index.layout",nil)
	
 }


	}

}
//Index page of the web application
func (lh *LoginHandler) Index(w http.ResponseWriter,r *http.Request, _ httprouter.Params){
	lh.tmpl.ExecuteTemplate(w,"index.layout",nil)
}
func (lh *LoginHandler) Companydashboard(w http.ResponseWriter,r *http.Request, _ httprouter.Params){
	lh.tmpl.ExecuteTemplate(w,"company_dashboard.layout",nil)
}
func (lh *LoginHandler) Admindashboard(w http.ResponseWriter,r *http.Request, _ httprouter.Params){
	lh.tmpl.ExecuteTemplate(w,"admindahsboard.layout",nil)
}
//chake the role
func (lh *LoginHandler) chakeRole(username string,password string) string{
	user  := []entity.User{}
	comps := []entity.Companie{}
	user,_ = lh.logService.Users()
	comps,_= lh.logService.Companies()
	var role string
	for _,u := range user{
		if u.Username == username && u.Password == password  {
		role = "user"
	}
}
for _,c := range comps{
	if c.FullName == username && c.Password == password{
		
		role = "company"
	}
}
return role
}
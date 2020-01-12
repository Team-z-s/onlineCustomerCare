package handler

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"gitlab.com/username/online-service-and-customer-care2.0/company/company_service"
	"gitlab.com/username/online-service-and-customer-care2.0/entity"
	"gitlab.com/username/online-service-and-customer-care2.0/service/service_service"
	"gitlab.com/username/online-service-and-customer-care2.0/user"
	"go/types"
	"html/template"
	"net/http"

)
type UserHandler struct {
	userService user.UserService
	tmpl        *template.Template
}
// NewUserHandler returns new AdminCommentHandler object
func NewUserHandler(userService user.UserService , T *template.Template) *UserHandler {
	return &UserHandler{userService: userService, tmpl: T}
}
func (uh *UserHandler) Signin(w http.ResponseWriter,r *http.Request, _ httprouter.Params){
	if r.Method == http.MethodPost {
		l := r.ContentLength
		body := make([]byte, l)
		r.Body.Read(body)
		usr := &entity.User{}
		usr.FName = r.FormValue("firstname")
		usr.LName = r.FormValue("lastname")
		usr.Email = r.FormValue("email")
		
		usr.Phone = r.FormValue("phone")
		usr.Username = r.FormValue("username")
		usr.Password = r.FormValue("password")

		err := json.Unmarshal(body, usr)

		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		usr, errs := uh.userService.StoreUser(usr)

		if len(errs) > 0 {
			w.Header().Set("Content-Type", "application/json")
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)

			return
		}
		http.Redirect(w, r, "/user/seatch", http.StatusSeeOther)

	} else {
		uh.tmpl.ExecuteTemplate(w, "signup.layout", nil)
	}


}
func (uh *UserHandler) Search(w http.ResponseWriter,r *http.Request, _ httprouter.Params){
	if r.Method == http.MethodPost{
		l := r.ContentLength
		body := make([]byte, l)
		r.Body.Read(body)
		serv := &entity.Service{}
		serv.Name= r.FormValue("search_form")

		err := json.Unmarshal(body, serv)

		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		serv,_ = service_service.NewServiceService().Service(serv.Name)
		company,_=company_service.NewCompanyService().Company(uint(serv.CompanyID))

		ran := make([]types.Object,2)
		ran = append(ran, serv)
		ran = append(ran,company)
		uh.tmpl.ExecuteTemplate(w, "browse.layout", ran)
	}


}

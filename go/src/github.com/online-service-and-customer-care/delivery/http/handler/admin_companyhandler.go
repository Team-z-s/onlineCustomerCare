package handler

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"gitlab.com/username/online-service-and-customer-care/entity"
	"gitlab.com/username/online-service-and-customer-care2.0/company"
	"html/template"
	"net/http"
	"strconv"
)
type AdminCompanyHandler struct {
	companyService company.CompanyService
	tmpl        *template.Template
}
// NewAdminCompanyHandler returns new AdminCommentHandler object
func NewAdminCompanyHandler(companyService company.CompanyService , T *template.Template) *AdminCompanyHandler {
	return &AdminCompanyHandler{companyService: companyService, tmpl: T}
}
func (ach *AdminCompanyHandler) AddCompany(w http.ResponseWriter,r *http.Request, _ httprouter.Params){
	if r.Method == http.MethodPost {
		l := r.ContentLength
		body := make([]byte, l)
		r.Body.Read(body)
		company := &entity.Companie{}
		company.FullName = r.FormValue("company_name")
		company.Email = r.FormValue("email")
		company.Address = r.FormValue("city") + r.FormValue("street_kebele")
		company.Phone = r.FormValue("phone")
		company.Moto = r.FormValue("moto")
		company.Password = r.FormValue("password")
		mf, fh, err := r.FormFile("comp_logo")
		if err != nil {
			panic(err)
		}
		defer mf.Close()

		company.Logo = fh.Filename

		writeFile(&mf, fh.Filename)

		err = json.Unmarshal(body, company)

		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		company, errs := ach.companyService.StoreCompany(company)

		if len(errs) > 0 {
			w.Header().Set("Content-Type", "application/json")
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)

			return
		}
		http.Redirect(w, r, "/admin_dashboard", http.StatusSeeOther)

	} else {
		ach.tmpl.ExecuteTemplate(w, "addcompany.layout", nil)
	}


}
func (ach *AdminCompanyHandler) DeleteCompany(w http.ResponseWriter,r *http.Request, _ httprouter.Params){
	if r.Method == http.MethodDelete{
		var id,_ = strconv.ParseUint(r.FormValue("id"),4,16)
		_,err := ach.companyService.DeleteCompany(uint(id))

		if len(err) > 0{
			w.Header().Set("Content-Type", "application/json")
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)

			return
		}
		http.Redirect(w, r, "/admin_dashboard", http.StatusSeeOther)

	}

}

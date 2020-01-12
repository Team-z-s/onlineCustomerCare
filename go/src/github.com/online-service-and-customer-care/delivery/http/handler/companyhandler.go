package handler
import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"gitlab.com/username/online-service-and-customer-care2.0/entity"
	"gitlab.com/username/online-service-and-customer-care2.0/company"

	"gitlab.com/username/online-service-and-customer-care2.0/service/service_repsitory"
	"gitlab.com/username/online-service-and-customer-care2.0/service/service_service"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)
type CompanyHandler struct {
	companyService company.CompanyService
	tmpl        *template.Template
}
// NewCompanyHandler returns new AdminCommentHandler object
func NewCompanyHandler(companyService company.CompanyService, T *template.Template) *CompanyHandler {
	return &CompanyHandler{companyService: companyService , tmpl: T}
}
func (ch *CompanyHandler) ChangeLogo(w http.ResponseWriter,r *http.Request, _ httprouter.Params){
	company := entity.Companie{}
	if r.Method == http.MethodPost{
		mf, fh, err := r.FormFile("comp_logo")
		if err != nil {
			panic(err)
		}
		defer mf.Close()

		company.Logo = fh.Filename

		writeFile(&mf, fh.Filename)

		_, errs := ch.companyService.UpdateCompany(&company)

		if len(errs) >0 {
			panic(err)
		}
		if len(errs) == 0{
			http.Redirect(w, r, "/company_dashboard", http.StatusSeeOther)
		}else{
			ch.tmpl.ExecuteTemplate(w,"changelogo.layout",nil)
		}


	}

}
func (ch *CompanyHandler) Addservice(w http.ResponseWriter,r *http.Request, _ httprouter.Params){
	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	service := entity.Service{}
	if r.Method == http.MethodPost{
		service.Name = r.FormValue("name")
		service.Description = r.FormValue("description")

		err := json.Unmarshal(body, service)

		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		servrepo := service_repsitory.NewServiceGormRepo(dbconn)
		service, errs := service_service.NewServiceService(servrepo).StoreService(&service)
		if len(errs) > 0 {
			w.Header().Set("Content-Type", "application/json")
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Redirect(w, r, "/company_dashboard", http.StatusSeeOther)

	} else {
		ch.tmpl.ExecuteTemplate(w, "addservice.layout", nil)

	}

}
func (ch *CompanyHandler) Take_attendance(w http.ResponseWriter,r *http.Request, _ httprouter.Params){

	/*/////////////////////////////////////////////////////////////////////////////////////////////////////////*/

}
func (ch *CompanyHandler) Calculate_salary(w http.ResponseWriter,r *http.Request, _ httprouter.Params){

	/*/////////////////////////////////////////////////////////////////////////////////////////////////////////*/

}
func writeFile(mf *multipart.File, fname string) {

	wd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	path := filepath.Join(wd, "ui", "assets", "img", fname)
	image, err := os.Create(path)

	if err != nil {
		panic(err)
	}
	defer image.Close()
	io.Copy(image, *mf)
}

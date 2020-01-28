package handler

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"onlineCustomerCare/Service"
	"onlineCustomerCare/entity"
	"onlineCustomerCare/login"
	"onlineCustomerCare/session"
	"os"
	"path/filepath"
	"strconv"

	//"onlineCustomerCare/entity"
	"onlineCustomerCare/company"
	"text/template"
)

type CompanyHandler struct {
	compservice company.CompanyService
	logservice login.LoginService
	tmpl *template.Template
	servService Service.ServiceService
	userSess   *entity.Session
}

func NewCompanyHandler(cs company.CompanyService, logser login.LoginService,T *template.Template )*CompanyHandler{
	return &CompanyHandler{compservice:cs,logservice:logser,tmpl:T}
}
type Detail struct {
	company *entity.Companie
	serv *entity.Service
}
func (ch *CompanyHandler) GetCompany(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet{
		ch.tmpl.ExecuteTemplate(w,"companydetail.html",nil)
	}else{
		cl := session.GetSessionData(w,r)
		comp,_ := ch.compservice.CompanyByName(cl.Username)
		serv,_ := ch.servService.ServiceById(comp.CompanyID)
		d := Detail{company:comp,serv:serv}
		ch.tmpl.ExecuteTemplate(w,"companydetail.html",d)
	}
	http.Redirect(w,r,"/user/search",http.StatusSeeOther)
}

func (ch *CompanyHandler) UpdateCompany(w http.ResponseWriter, r *http.Request){
	comp := entity.Companie{}
	account := entity.Account{}

	if r.Method == http.MethodGet{
		ch.tmpl.ExecuteTemplate(w,"update_company.html",nil)
	}else{
		comp.CompanyID = retutnid(ch) + 1
		comp.FullName = r.FormValue("Company_name")
		comp.Email = r.FormValue("Email")
		comp.Address = r.FormValue("Address")
		comp.Moto = r.FormValue("Moto")
		comp.Phone = r.FormValue("Phone")
		comp.Password = r.FormValue("Password")
		//comp.Logo = r.FormValue("Logo")
		//############################################################################################################33
		mf, fh, err := r.FormFile("Logo")

		//fmt.Println(mf)
		//fmt.Println(fh)

		if err != nil {
			panic(err)
		}
		defer mf.Close()

		comp.Logo = fh.Filename

		writeFile(&mf, fh.Filename)

		account.Username = comp.FullName
		account.Password = comp.Password
		account.Role_id = 2

		//ch.logservice.StoreAccount(&account)
		_,errs := ch.compservice.UpdateCompany(&comp)
		if len(errs) >0 {
			panic(errs)
			ch.tmpl.ExecuteTemplate(w,"update_company.html",nil)

		}
		http.Redirect(w,r,"/admin_dashboard",http.StatusSeeOther)
	}
	}


func(ch *CompanyHandler) DeleteCompany(w http.ResponseWriter , r * http.Request){
	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		}


		_, errs := ch.compservice.DeleteCompany(uint(id))
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}
	}
	http.Redirect(w, r, "/admin_dashboard", http.StatusSeeOther)
}

func (ch *CompanyHandler) AddCompany(w http.ResponseWriter, r *http.Request){
	comp := entity.Companie{}
	account := entity.Account{}
	var cookie, errss = r.Cookie("session")
	if errss == nil {
		cookievalue1 := cookie.Value
		fmt.Println(cookievalue1)
	}

	if r.Method == http.MethodGet{
		ch.tmpl.ExecuteTemplate(w,"add_company1.html",nil)
	}else{
		//fmt.Println(ch.userSess.Uername)
		comp.CompanyID = retutnid(ch) + 1
		comp.FullName = r.FormValue("Company_name")
		comp.Email = r.FormValue("Email")
		comp.Address = r.FormValue("Address")
		comp.Moto = r.FormValue("Moto")
		comp.Phone = r.FormValue("Phone")
		comp.Password = r.FormValue("Password")
		//comp.Logo = r.FormValue("Logo")
//############################################################################################################33
		mf, fh, err := r.FormFile("Logo")

		//fmt.Println(mf)
		//fmt.Println(fh)

		if err != nil {
			panic(err)
		}
		defer mf.Close()

		comp.Logo = fh.Filename

		writeFile(&mf, fh.Filename)

		account.Username = comp.FullName
		account.Password = comp.Password
		account.Role_id = 2

		ch.logservice.StoreAccount(&account)
		_,errs := ch.compservice.StoreCompany(&comp)
		if len(errs) >0 {
			panic(errs)
			ch.tmpl.ExecuteTemplate(w,"add_company1.html",nil)

		}
		http.Redirect(w,r,"/admin_dashboard",http.StatusSeeOther)
	}

}
func (ch *CompanyHandler)	AdminDashboard(w http.ResponseWriter, r *http.Request){
	comp := []entity.Companie{}
	comp ,_ = ch.compservice.Companies()
	if r.Method == http.MethodGet {
		fmt.Println(comp)
		if len(comp) == 0 {
			ch.tmpl.ExecuteTemplate(w, "search.html", nil)
		} else {
			ch.tmpl.ExecuteTemplate(w, "search.html", comp)
		}
	}
	count := 0
	cc := entity.Companie{}
	if r.Method == http.MethodPost{
		name := r.FormValue("keyword")
		for _,c := range comp{
			if c.FullName == name{
				count = count + 1
				cc = c
			}
		}
		fmt.Println(cc)
		if count == 0{
			ch.tmpl.ExecuteTemplate(w, "search.html", nil)
		}else{
			fmt.Println("######################################")
			ch.tmpl.ExecuteTemplate(w, "search.html", cc)
		}
	}


}
func (ch *CompanyHandler) Search(w http.ResponseWriter,r *http.Request) {
	comp := []entity.Companie{}
	comp, _ = ch.compservice.Companies()
	if r.Method == http.MethodGet {
		fmt.Println(comp)
		if len(comp) == 0 {
			ch.tmpl.ExecuteTemplate(w, "browse.html", nil)
		} else {
			ch.tmpl.ExecuteTemplate(w, "browse.html", comp)
		}
	}
	count := 0
	cc := entity.Companie{}
	if r.Method == http.MethodPost {
		name := r.FormValue("keyword")
		for _, c := range comp {
			if c.FullName == name {
				count = count + 1
				cc = c
			}
		}
		fmt.Println(cc)
		if count == 0 {
			ch.tmpl.ExecuteTemplate(w, "browse.html", nil)
		} else {
			fmt.Println("######################################")
			ch.tmpl.ExecuteTemplate(w, "browse.html", cc)
		}
		ch.tmpl.ExecuteTemplate(w, "browse.html", nil)
	}
}
func retutnid(ch *CompanyHandler) int {
	comp, _ := ch.compservice.Companies()
	var i int = comp[0].CompanyID
	for _, c:= range comp{
		if i < c.CompanyID {
			i = c.CompanyID
		}
	}
	return i
}

func writeFile(mf *multipart.File, fname string) {

	wd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	path := filepath.Join(wd, "../../ui", "assets", "images", fname)
	//fmt.Println(path)
	image, err := os.Create(path)

	if err != nil {
		panic(err)
	}
	defer image.Close()
	io.Copy(image, *mf)
}








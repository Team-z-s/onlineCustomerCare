package handler

import (
	"net/http"
	"onlineCustomerCare/Service"
	"onlineCustomerCare/company"
	"onlineCustomerCare/entity"
	"onlineCustomerCare/session"
	"text/template"
)

type ServiceHandler struct{
	servService Service.ServiceService
	compService company.CompanyService
	temp *template.Template
}
func NewServiceHandler(serv Service.ServiceService,cs company.CompanyService,t *template.Template)*ServiceHandler{
	return &ServiceHandler{servService:serv,compService:cs,temp:t}
}
func(sh *ServiceHandler) AddService(w http.ResponseWriter, r *http.Request){
	serv := entity.Service{}
	cl := session.GetSessionData(w,r)
	if r.Method == http.MethodGet{
		sh.temp.ExecuteTemplate(w,"add_service.html",nil)
	}else{
		comp,errcc := sh.compService.CompanyByName(cl.Username)
		if len(errcc)>0{
			return
		}else{
			serv.CompanyID = comp.CompanyID
			serv.ID = retutnservid(sh) + 1
			serv.Name = r.FormValue("Name")
			serv.Description = r.FormValue("Description")

			_,errs := sh.servService.StoreService(&serv)
			if len(errs) > 0{
				sh.temp.ExecuteTemplate(w,"add_service.html",nil)
			}
			http.Redirect(w,r,"/company_dashbord",http.StatusSeeOther)

		}

	}



}


func(sh *ServiceHandler) DeleteService(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("name")
		//id, err := strconv.Atoi(idRaw)
		_,err := sh.servService.DeleteService(idRaw)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		}
		_, errs := sh.servService.DeleteService(idRaw)
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}
	}
	http.Redirect(w, r, "/admin_dashboard", http.StatusSeeOther)
}

func (sh *ServiceHandler) UpdateService(w http.ResponseWriter, r *http.Request){
	serv := entity.Service{}
	cl := session.GetSessionData(w,r)
	if r.Method == http.MethodGet{
		sh.temp.ExecuteTemplate(w,"add_service.html",nil)
		}
	if r.Method == http.MethodPost{
		comp,_ := sh.compService.CompanyByName(cl.Username)
		serv.CompanyID = comp.CompanyID

		serv.Name = r.FormValue("Name")
		serv.Description = r.FormValue("Description")

		_,errs := sh.servService.UpdateService(&serv)
		if len(errs) > 0{
			sh.temp.ExecuteTemplate(w,"add_service.html",nil)
		}
		http.Redirect(w,r,"/company_dashbord",http.StatusSeeOther)
	}
}



func retutnservid(sh *ServiceHandler) int {
	var i int
	serv, _ := sh.servService.Services()
	if len(serv)==0{
		return 0
	}else{
		i = serv[0].ID
		for _, s:= range serv{
			if i < s.ID {
				i = s.ID
			}
		}
	}
	return i
}
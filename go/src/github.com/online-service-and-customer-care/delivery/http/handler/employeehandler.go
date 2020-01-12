package handler

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"gitlab.com/username/online-service-and-customer-care2.0/employee"
	"gitlab.com/username/online-service-and-customer-care2.0/entity"
)

// EmployeeHandler handles comment related http requests
type EmployeeHandler struct {
	emplService employee.EmployeeService
	tmpl        *template.Template
}

// NewEmployeeHandler returns new AdminCommentHandler object
func NewEmployeeHandler(emplService employee.EmployeeService, T *template.Template) *EmployeeHandler {
	return &EmployeeHandler{emplService: emplService, tmpl:T}
}

// GetEmployees handles GET /v1/admin/comments request
func (eh *EmployeeHandler) GetEmployees(w http.ResponseWriter,
	r *http.Request, _ httprouter.Params) {

	employee, errs := eh.emplService.Employees()

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(employee, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}

// GetSingleEmployee handles GET /v1/admin/comments/:id request
func (eh *EmployeeHandler) GetSingleEmployee(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	employee, errs := eh.emplService.Employee(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(employee, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

//Index dsng

//Add fjdgl
func (eh *EmployeeHandler) Add(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	eh.tmpl = template.Must(template.ParseGlob("../../ui/templates/*"))
	eh.tmpl.ExecuteTemplate(w, "addemployee.layout", nil)
}

// PostEmployee handles POST /v1/admin/comments request
func (eh *EmployeeHandler) PostEmployee(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Method == http.MethodPost {
		l := r.ContentLength
		body := make([]byte, l)
		r.Body.Read(body)
		employee := &entity.Employee{}
		employee.FName = r.FormValue("firstname")
		employee.LName = r.FormValue("lastname")
		employee.Email = r.FormValue("email")
		employee.Address = r.FormValue("city") + r.FormValue("street_kebele")
		employee.Phone = r.FormValue("phone")
		employee.Sallary, _ = strconv.ParseFloat(r.FormValue("salary"), 64)
		employee.Username = r.FormValue("username")
		employee.Password = r.FormValue("password")

		err := json.Unmarshal(body, employee)

		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		employee, errs := eh.emplService.StoreEmployee(employee)

		if len(errs) > 0 {
			w.Header().Set("Content-Type", "application/json")
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)

			return
		}
		http.Redirect(w, r, "/Index", http.StatusSeeOther)

	} else {
		eh.tmpl.ExecuteTemplate(w, "addemployee.layout", nil)
	}

}

// DeleteEmployee handles DELETE /v1/admin/comments/:id request
func (eh *EmployeeHandler) DeleteEmployee(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	_, errs := eh.emplService.DeleteEmployee(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	return
}

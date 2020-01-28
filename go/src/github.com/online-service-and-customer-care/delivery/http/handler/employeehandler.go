package handler

import (

	"net/http"
	"onlineCustomerCare/company"
	"onlineCustomerCare/employee"
	"onlineCustomerCare/entity"
	"onlineCustomerCare/login"
	"onlineCustomerCare/session"
	"strconv"
	"text/template"
)

type EmployeeHandler struct {
	emplService employee.EmployeeService
	logservice login.LoginService
	compService company.CompanyService
	temp *template.Template
}

func NewEmloyeeHandler(es employee.EmployeeService,log login.LoginService,cs company.CompanyService,T *template.Template)*EmployeeHandler{
	return &EmployeeHandler{emplService:es,logservice:log,compService:cs,temp:T}
}
func (eh *EmployeeHandler) AddEmployee(w http.ResponseWriter, r *http.Request){
	empl := entity.Employee{}
	act := entity.Account{}
	cl := session.GetSessionData(w,r)

	if r.Method == http.MethodGet {
		eh.temp.ExecuteTemplate(w, "add_employee.html", nil)
	}else {
		comp,errcc := eh.compService.CompanyByName(cl.Username)
		if len(errcc)>0{

		}else {
			empl.CompanyID = comp.CompanyID
			empl.EmployeeID = retutnemployeeid(eh)
			empl.FName = r.FormValue("FName")
			empl.LName = r.FormValue("LName")
			empl.Email = r.FormValue("Email")
			empl.Phone = r.FormValue("Phone")
			empl.Address = r.FormValue("Address")
			empl.Username = r.FormValue("Username")
			empl.Password = r.FormValue("password")
			empl.Salary, _ = strconv.ParseFloat(r.FormValue("Salary"), 64)

			mf, fh, err := r.FormFile("Photo")

			//fmt.Println(mf)
			//fmt.Println(fh)

			if err != nil {
				panic(err)
			}
			defer mf.Close()

			empl.Phone = fh.Filename

			writeFile(&mf, fh.Filename)

			act.Username = r.FormValue("Username")
			act.Password = r.FormValue("password")
			act.Role_id = 3

			_, errs := eh.emplService.StoreEmployee(&empl)
			_, errsa := eh.logservice.StoreAccount(&act)
			if len(errs) > 0 || len(errsa) > 0 {
				eh.temp.ExecuteTemplate(w, "add_employee.html", nil)
			}
			http.Redirect(w, r, "/company_dashboard", http.StatusSeeOther)
		}
	}
}
func (eh *EmployeeHandler) DeleteEmployee(w http.ResponseWriter, r *http.Request){
	id,_ := strconv.Atoi( r.FormValue("ID"))
	_,err := eh.emplService.DeleteEmployee(uint(id))
	if len(err) > 0{
		eh.temp.ExecuteTemplate(w, "delete_employee.html", nil)
	}
	http.Redirect(w,r,"/##########################",http.StatusSeeOther)
}

func (eh *EmployeeHandler) AssignTask(w http.ResponseWriter, r *http.Request){

}
func (eh *EmployeeHandler) ShowProfile(w http.ResponseWriter, r *http.Request){
	//####################################### find the data of the employee from the session and execute the template with the provided data


	eh.temp.ExecuteTemplate(w,"employee_profile.html",nil)
}

func retutnemployeeid(eh *EmployeeHandler) int {
	var i int
	empl, _ := eh.emplService.Employees()
	if len(empl) == 0{
		return 0
	}else {
		i = empl[0].EmployeeID
		for _, e := range empl {
			if i < e.EmployeeID {
				i = e.EmployeeID
			}
		}
	}
	return i
}




package handler

import (
	//"context"
	//"onlineCustomerCare/rtoken"
	//"onlineCustomerCare/validation"

	"time"

	//"fmt"
	//"github.com/dgrijalva/jwt-go"
	"net/http"
	"onlineCustomerCare/entity"
	"onlineCustomerCare/login"
	//"onlineCustomerCare/rtoken"
	"onlineCustomerCare/session"
	//"onlineCustomerCare/validation"
	"text/template"
	//"time"
)

type LoginHandler struct {
	logservice login.LoginService
	temp *template.Template

	loggedInUser   *entity.User
}


func NewLoginHandler(logserv login.LoginService, T *template.Template)*LoginHandler {
	return &LoginHandler{logservice: logserv, temp: T}
}
func (lh *LoginHandler) Login(w http.ResponseWriter, r *http.Request) {
	act := entity.Account{}
	if r.Method == http.MethodGet {
		lh.temp.ExecuteTemplate(w, "login.html", nil)

	} else {
			act.Username =r.FormValue("username")
			act.Password = r.FormValue("password")
			account, err := lh.logservice.Account(act.Username)
			tokenString, err := session.Generate(act.Username)
			if err != nil {
				return
			}
			http.SetCookie(w, &http.Cookie{
				Name:  "token",
				Value: tokenString,
			})
			//fmt.Println(tokenString)
			//cl := session.GetSessionData(w,r)
			//fmt.Println(cl.Username)
			Verrors := "Invalid Username or Password"
			if err != nil {
				lh.temp.ExecuteTemplate(w, "login.html", Verrors)
			}else{
				if len(Verrors) >0 || act.Password == account.Password {
					role := account.Role_id
						switch role {
							case 1:

								http.Redirect(w, r, "/admin_dashboard", http.StatusSeeOther)

								break
							case 2:
								http.Redirect(w, r, "/company_dashboard", http.StatusSeeOther)
								break
							case 3:
								http.Redirect(w, r, "/employee/profile", http.StatusSeeOther)
								break
							case 4:
								http.Redirect(w, r, "/user/search", http.StatusSeeOther)
								break
							default:
								lh.temp.ExecuteTemplate(w, "login.html", Verrors)
						}
					}else{
						lh.temp.ExecuteTemplate(w, "login.html", Verrors)
					}
				}
	}
}


func (lh *LoginHandler) Logout(w http.ResponseWriter, r *http.Request) {

	c := http.Cookie{
		Name:    "token",
		MaxAge:  -1,
		Expires: time.Unix(1, 0),
		Value:   "",
	}
	http.SetCookie(w, &c)
	http.Redirect(w, r, "/Index", http.StatusSeeOther)
}

func (lh *LoginHandler) CompanyDashboard(w http.ResponseWriter, r *http.Request){
	lh.temp.ExecuteTemplate(w,"company_dashboard.html",nil)
}


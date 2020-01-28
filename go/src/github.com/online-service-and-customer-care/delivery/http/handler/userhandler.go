package handler

import (
	"fmt"
	"net/http"
)
import (
	"onlineCustomerCare/entity"
	"onlineCustomerCare/user"
	"text/template"
)

type UserHandler struct{
	userService user.UserService
	tmpl *template.Template

}

func NewUserHandler (userServ user.UserService,tmpl *template.Template)*UserHandler{
	return &UserHandler{userService:userServ,tmpl:tmpl}
}

func (uh *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	user := entity.User{}
	if r.Method == "GET" {
		uh.tmpl.ExecuteTemplate(w,"signup.html",nil)

	} else {
		user.ID = retutnuid(uh) + 1
		user.FName = r.FormValue("first_name")
		user.LName = r.FormValue("last_name")
		user.Email = r.FormValue("email")
		user.Phone = r.FormValue("phone")
		user.Username = r.FormValue("username")
		user.Password = r.FormValue("password")

		_, err := uh.userService.StoreUser(&user)
		if len(err) > 0 {
			fmt.Println("error")
			uh.tmpl.ExecuteTemplate(w, "signup.html", nil)
		}
		fmt.Println("success")
		http.Redirect(w, r, "/user/search", http.StatusSeeOther)

	}
}

func retutnuid(uh *UserHandler) int {
	users,_ := uh.userService.Users()
	var i int = users[0].ID
	for _, u:= range users{
		if i < u.ID {
			i = u.ID
		}
	}
	return i
}

package apiserver

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("pkg/static/templates/*"))
}

func (c *APIServer) Register(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/register" {
		return
	}

	switch r.Method {
	case http.MethodGet:
		tpl.ExecuteTemplate(w, "register.html", nil)
	case http.MethodPost:
		r.ParseForm()
		userName := r.FormValue("name")
		userEmail := r.FormValue("email")
		userPassword := r.FormValue("Password")
		fmt.Println(userEmail, userName, userPassword, "asda")
	}
}

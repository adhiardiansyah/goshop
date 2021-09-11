package controllers

import (
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
	// "github.com/unrolled/render"
)

var tpl *template.Template
var login = sessions.NewCookieStore([]byte("mysession"))

func ParseTemplates() (*template.Template, error) {

	if t, _ := tpl.ParseGlob("templates/*.html"); t != nil {
		tpl = t
	}
	if t, _ := tpl.ParseGlob("templates/pages/*.html"); t != nil {
		tpl = t
	}
	return tpl.ParseGlob("/*.html")
}

func Home(w http.ResponseWriter, r *http.Request) {
	// render := render.New(render.Options{
	// 	Layout: "layout",
	// })

	// _ = render.HTML(w, http.StatusOK, "home", map[string]interface{}{
	// 	"title": "Home Title",
	// 	"body":  "Home Description",
	// })

	// tpl, _ = tpl.ParseGlob("templates/*.html")
	ParseTemplates()
	tpl.ExecuteTemplate(w, "home.html", nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	ParseTemplates()
	tpl.ExecuteTemplate(w, "login.html", nil)
}

func ProcessLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	if username == "abc" && password == "123" {
		session, _ := login.Get(r, "mysession")
		session.Values["username"] = username
		session.Save(r, w)
		http.Redirect(w, r, "/welcome", http.StatusSeeOther)
	} else {
		data := map[string]interface{}{
			"err": "Invalid",
		}
		ParseTemplates()
		tpl.ExecuteTemplate(w, "login.html", data)
	}
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	session, _ := login.Get(r, "mysession")
	username := session.Values["username"]
	data := map[string]interface{}{
		"username": username,
	}
	ParseTemplates()
	tpl.ExecuteTemplate(w, "welcome.html", data)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := login.Get(r, "mysession")
	session.Options.MaxAge = -1
	session.Save(r, w)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

package controllers

import (
	"html/template"
	"net/http"
	// "github.com/unrolled/render"
)

var tpl *template.Template

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

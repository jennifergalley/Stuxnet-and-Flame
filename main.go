package main

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	r := httprouter.New()
	http.Handle("/", r)
	r.GET("/", index)
	r.GET("/:page", serve)
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public/"))))

	tpl = template.New("roottemplate")
	tpl = template.Must(tpl.ParseGlob("templates/*.html"))
}

func index(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	tpl.ExecuteTemplate(res, "index.html", nil)
}

func serve(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	page := ps.ByName("page")
	tpl.ExecuteTemplate(res, page + ".html", nil)
}

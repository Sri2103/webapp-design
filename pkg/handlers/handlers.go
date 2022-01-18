package handlers

import (
	"myapp/pkg/renders"
	"net/http"
)

func HelloServer(w http.ResponseWriter, r *http.Request) {

	//fmt.Fprintf(w, "hello World")
	w.Write([]byte("<h1>HelloWorld</h1>"))
}

//  Home page Handler
func HomePage(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, "home.page.html")
	//Tpl.Execute(w, nil)
}

//	2. About Page handler
func About(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, "about.page.html")
	//Ab.Execute(w, nil)s
}

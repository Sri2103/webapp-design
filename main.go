package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

// using  newmux
// create a (server function/ Handler) to handle request and write response
func HelloServer(w http.ResponseWriter, r *http.Request) {

	//fmt.Fprintf(w, "hello World")
	w.Write([]byte("<h1>HelloWorld</h1>"))
}

// Home Page Template
var tpl = template.Must(template.ParseFiles("./templates/home.html"))

//  Home page Handler
func HomePage(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

//	2. About Page handler
var ab = template.Must(template.ParseFiles("./templates/about.html"))

func About(w http.ResponseWriter, r *http.Request) {
	ab.Execute(w, nil)
}

func main() {
	mux := http.NewServeMux()

	//Render the Pages templates

	//handle the server function with mux router

	mux.HandleFunc("/", HelloServer)

	mux.HandleFunc("/home", HomePage)

	mux.HandleFunc("/about", About)

	//getting Port value from env file
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// address port to handle above server
	log.Printf("listening on Port :%s", port)

	http.ListenAndServe(":"+port, mux)

}

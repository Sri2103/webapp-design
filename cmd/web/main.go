package main

import (
	"log"
	"myapp/pkg/config"
	"myapp/pkg/handlers"
	"myapp/pkg/renders"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
)

// create a (server function/ Handler) to handle request and write response
// func HelloServer(w http.ResponseWriter, r *http.Request) {

// 	//fmt.Fprintf(w, "hello World")
// 	w.Write([]byte("<h1>HelloWorld</h1>"))
// }

var app config.AppConfig

var session *scs.SessionManager

func main() {
	// http.HandleFunc("/hello", HelloServer)
	//mux := http.NewServeMux()

	//Render the Pages templates

	//handle the server function with mux router

	// mux.HandleFunc("/home", handlers.HomePage)

	// mux.HandleFunc("/about", handlers.About)

	//getting Port value from env file

	app.Inproduction = false

	session = scs.New()

	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.Inproduction

	app.Session = session

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	renders.NewTemplates(&app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("listening on Port :%s", port)

	srv := http.Server{
		Addr:    ":" + port,
		Handler: routes(&app),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	// address port to handle above server

}

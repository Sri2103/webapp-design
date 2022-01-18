package handlers

import (
	"myapp/pkg/config"
	"myapp/pkg/models"
	"myapp/pkg/renders"
	"net/http"
)

// Create a Respository struct which contains appconfig as object

var Repo *Respository

type Respository struct {
	App *config.AppConfig
}

// Create a function that utilizes appconfig to form a Repo pointer to struct Repository and therby New Repository

func NewRepo(a *config.AppConfig) *Respository {
	return &Respository{
		App: a,
	}
}

func NewHandlers(r *Respository) {
	Repo = r
}

func HelloServer(w http.ResponseWriter, r *http.Request) {

	//fmt.Fprintf(w, "hello World")
	w.Write([]byte("<h1>HelloWorld</h1>"))
}

//  Home page Handler
// func HomePage(w http.ResponseWriter, r *http.Request) {
// 	renders.RenderTemplate(w, "home.page.html")
// }

func (m *Respository) Home(w http.ResponseWriter, r *http.Request) {
	remotIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remotIp", remotIP)
	renders.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

//	2. About Page handler
func (m *Respository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello Again"

	remoteIp := m.App.Session.GetString(r.Context(), "remote_Ip")

	stringMap["remote_Ip"] = remoteIp

	renders.RenderTemplate(w, "about.Page.html", &models.TemplateData{
		StringMap: stringMap,
	})

	// renders.RenderTemplate(w, "about.page.html",&models.TemplateData{})

}

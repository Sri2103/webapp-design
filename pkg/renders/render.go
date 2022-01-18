package renders

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"myapp/pkg/config"
	"myapp/pkg/models"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

//defining app configuration from config to use for rendering

var app *config.AppConfig

// creating a new template with Appconfig  as base

func NewTemplates(a *config.AppConfig) {
	app = a
}

// introducing dataStructure  of template from models as td

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, filename string, td *models.TemplateData) {
	var tc map[string]*template.Template

	if app.Usecache {
		//if u can useCache stored in TemplateCache get templateCache from AppConfig otherwise Create new Cache from Createcache()
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	//create a cache initiallly to render them step by step
	// tc, err := CreateTemplateCache()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	t, ok := tc[filename]
	log.Print(t)
	if !ok {
		log.Fatal("Could not get template from Cache")
	}

	// sending data to browser in the form of bytes
	buf := new(bytes.Buffer)
	//executing string template retreived from the map of templates
	//_ = t.Execute(buf, nil)

	// Add data to Buffer before sending to Browser
	td = AddDefaultData(td)
	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error occured while writing Template to Browser")
	}

}

// create a template Cache : everytime a a request runs cache is formed to be  renderered by RenderTemplate Function
func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")

	if err != nil {
		return myCache, err
	}

	for _, page := range pages {

		name := filepath.Base(page)
		// Parsing the results from name with tmpl extension. As base gives out extensions
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts

	}

	log.Print(myCache)

	return myCache, err

}

package renders

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// var Tpl = template.Must(template.ParseFiles("./templates/home.html"))
// var Ab = template.Must(template.ParseFiles("./templates/about.html"))

var functions = template.FuncMap{}

func RenderTemplate(w http.ResponseWriter, filename string) {

	//create a cache initiallly to render them step by step
	tc, err := CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	t, ok := tc[filename]
	log.Print(t)
	if !ok {
		log.Fatal("Could not get template from Cache")
	}

	// sending data to browser in the form of bytes
	buf := new(bytes.Buffer)
	//executing string template retreived from the map of templates
	_ = t.Execute(buf, nil)
	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error occured while writing Template to Browser")
	}

}

// create a template Cache : everytime a a request runs cache is formed to be  renderered by RenderTemplate Function
func CreateTemplateCache() (map[string]*template.Template, error) {

	templateCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")

	if err != nil {
		return templateCache, err
	}

	for _, page := range pages {

		name := filepath.Base(page)
		// Parsing the results from name with tmpl extension. As base gives out extensions
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return templateCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return templateCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return templateCache, err
			}
		}

		templateCache[name] = ts

	}

	log.Print(templateCache)

	return templateCache, err

}

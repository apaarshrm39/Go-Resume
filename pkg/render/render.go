package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/apaarshrm39/Go-Resume/pkg/config"
	"github.com/apaarshrm39/Go-Resume/pkg/models"
)

var functions = template.FuncMap{}

var appSettings *config.Appconfig

func SetAppsettings(a *config.Appconfig) {
	appSettings = a
}

func Render(w http.ResponseWriter, tmpl string, tdata *models.TemplateData) {

	templates := make(map[string]*template.Template)
	if appSettings.EnableCache {
		templates = appSettings.Templates
	} else {
		templates = TemplateCache()
	}
	templates[tmpl].Execute(w, tdata)
}

func TemplateCache() map[string]*template.Template {

	cache := make(map[string]*template.Template)
	files, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		name := filepath.Base(file)
		fmt.Println("working on file", name)

		tc, err := template.New(name).Funcs(functions).ParseFiles(file)
		if err != nil {
			log.Fatal(err)
		}

		layouts, err := filepath.Glob("./templates/*.layout.html")
		fmt.Println("size of layout", layouts)
		if err != nil {
			log.Fatal(err)
		}

		if len(layouts) > 0 {
			tc, err = tc.ParseGlob("./templates/*.layout.html")
			if err != nil {
				log.Fatal(err)
			}
		}
		cache[name] = tc
	}

	return cache
}

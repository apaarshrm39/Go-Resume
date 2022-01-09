package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

func Render(w http.ResponseWriter, tmpl string) {
	/*
		t, err := template.New("home.page.html").ParseFiles("./templates/" + tmpl)
		if err != nil {
			fmt.Println(err)
			return
		}
	*/
	//buf := new(bytes.Buffer)
	templates := TemplateCache()
	templates[tmpl].Execute(w, nil)
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

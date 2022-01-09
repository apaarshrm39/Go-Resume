package main

import (
	"log"
	"net/http"

	"github.com/apaarshrm39/Go-Resume/pkg/config"
	"github.com/apaarshrm39/Go-Resume/pkg/handlers"
	"github.com/apaarshrm39/Go-Resume/pkg/render"
)

func main() {

	mux := http.NewServeMux()

	dev_mode := config.Appconfig{
		EnableCache: false,
		Templates:   render.TemplateCache(),
	}

	render.SetAppsettings(&dev_mode)
	repository := handlers.NewRepo(&dev_mode)

	mux.HandleFunc("/", repository.Home)
	log.Fatal(http.ListenAndServe(":9000", mux))
}

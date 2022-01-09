package handlers

import (
	"net/http"

	"github.com/apaarshrm39/Go-Resume/pkg/config"
	"github.com/apaarshrm39/Go-Resume/pkg/models"
	"github.com/apaarshrm39/Go-Resume/pkg/render"
)

// Implimenting the Repository model

type Repository struct {
	AppSettings config.Appconfig
}

// Global variable to store app settings
var Repovar Repository

// Function to create a new Repo

func NewRepo(app *config.Appconfig) *Repository {
	return &Repository{
		AppSettings: *app,
	}
}

// function to set Repository for handler

func SetHandler(r *Repository) {
	Repovar = *r
}

func (repo *Repository) Home(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]string)
	data["dev_mode"] = "No chache"
	render.Render(w, "home.page.html", &models.TemplateData{
		StringMap: data,
	})
}

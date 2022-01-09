package handlers

import (
	"net/http"

	"github.com/apaarshrm39/Go-Resume/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.Render(w, "home.page.html")
}

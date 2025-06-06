package handlers

import (
	"net/http"

	"github.com/PradeepPadmanabanC/golang_learning/learn_html_templates/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}

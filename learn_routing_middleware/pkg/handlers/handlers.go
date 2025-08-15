package handlers

import (
	"net/http"

	"github.com/PradeepPadmanabanC/golang_learning/learn_html_templates/pkg/config"
	"github.com/PradeepPadmanabanC/golang_learning/learn_html_templates/pkg/models"
	"github.com/PradeepPadmanabanC/golang_learning/learn_html_templates/pkg/render"
)

var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// Perform logic

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

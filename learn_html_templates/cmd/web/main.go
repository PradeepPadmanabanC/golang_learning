package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PradeepPadmanabanC/golang_learning/learn_html_templates/pkg/config"
	"github.com/PradeepPadmanabanC/golang_learning/learn_html_templates/pkg/handlers"
	"github.com/PradeepPadmanabanC/golang_learning/learn_html_templates/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCahe()
	if err != nil {
		log.Fatal("cannot create template cache")

	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", repo.Home)
	http.HandleFunc("/about", repo.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}

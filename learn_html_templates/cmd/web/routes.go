package main

import (
	"net/http"

	"github.com/PradeepPadmanabanC/golang_learning/learn_html_templates/pkg/config"
	"github.com/PradeepPadmanabanC/golang_learning/learn_html_templates/pkg/handlers"
	"github.com/bmizerany/pat"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}

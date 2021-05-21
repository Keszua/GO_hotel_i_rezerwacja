package main

import (
	"net/http"

	"example.com/udemy/cmd/pkg/config"
	"example.com/udemy/cmd/pkg/handlers"
	"github.com/bmizerany/pat"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/", http.HandlerFunc(handlers.Repo.About))

	return mux
}

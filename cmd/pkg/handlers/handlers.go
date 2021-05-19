package handlers

import (
	"net/http"

	"example.com/udemy/cmd/pkg/render"
	//"example.com/cmd/pkg/render"
	//"example.com/udemy/cmd/pkg/render"
)

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.html")
}

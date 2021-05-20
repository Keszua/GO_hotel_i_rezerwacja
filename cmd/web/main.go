package main

import (
	"fmt"
	"log"
	"net/http"

	"example.com/udemy/cmd/pkg/config"
	"example.com/udemy/cmd/pkg/handlers"
	"example.com/udemy/cmd/pkg/render"
)

const portNumber = ":8080"

//main is the main appliction function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	render.NewTemplates(&app)

	fmt.Println("Starting aplication on port", portNumber)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	_ = http.ListenAndServe(portNumber, nil)

}

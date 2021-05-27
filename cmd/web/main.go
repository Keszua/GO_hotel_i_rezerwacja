package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/keszua/hotel/pkg/config"
	"github.com/keszua/hotel/pkg/handlers"
	"github.com/keszua/hotel/pkg/render"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

//main is the main appliction function
func main() {

	//change this to ttrue when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour //ważność sesji na 24 godziny
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteDefaultMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	fmt.Println("Starting aplication on port", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/burhanuday/basic-web-application/pkg/config"
	"github.com/burhanuday/basic-web-application/pkg/handlers"
	"github.com/burhanuday/basic-web-application/pkg/render"
)

const port = ":5000"

var app config.AppConfig
var session *scs.SessionManager

// main is the entrypoint
func main() {

	// change to true when Prod
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("Could not create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false
	app.Session = session

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()

	if err != nil {
		fmt.Println("error while starting server")
	}
}

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/burhanuday/basic-web-application/pkg/config"
	"github.com/burhanuday/basic-web-application/pkg/handlers"
	"github.com/burhanuday/basic-web-application/pkg/render"
)

const port = ":5000"

// main is the entrypoint
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("Could not create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

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

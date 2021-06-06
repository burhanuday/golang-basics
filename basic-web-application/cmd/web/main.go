package main

import (
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

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	http.ListenAndServe(port, nil)
}

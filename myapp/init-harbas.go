package main

import (
	"github.com/shreejilucifer/harbas"
	"log"
	"myapp/handlers"
	"os"
)

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	hbs := &harbas.Harbas{}
	err = hbs.New(path)
	if err != nil {
		log.Fatal(err)
	}

	hbs.AppName = "myapp"
	hbs.Debug = true

	myHandlers := &handlers.Handlers{
		App: hbs,
	}

	app := &application{
		App:      hbs,
		Handlers: myHandlers,
	}

	app.App.Routes = app.routes()

	return app
}

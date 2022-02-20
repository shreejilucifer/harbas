package main

import (
	"github.com/shreejilucifer/harbas"
	"log"
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

	hbs.InfoLog.Println("Debug is set to", hbs.Debug)

	app := &application{
		App: hbs,
	}

	return app
}

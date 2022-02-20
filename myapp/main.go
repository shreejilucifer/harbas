package main

import (
	"github.com/shreejilucifer/harbas"
	"myapp/handlers"
)

type application struct {
	App      *harbas.Harbas
	Handlers *handlers.Handlers
}

func main() {
	h := initApplication()
	h.App.ListenAndServe()
}

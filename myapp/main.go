package main

import "github.com/shreejilucifer/harbas"

type application struct {
	App *harbas.Harbas
}

func main() {
	h := initApplication()
	h.App.ListenAndServe()
}

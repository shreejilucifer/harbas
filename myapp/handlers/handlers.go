package handlers

import (
	"github.com/shreejilucifer/harbas"
	"net/http"
)

type Handlers struct {
	App *harbas.Harbas
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	err := h.App.Render.Page(w, r, "home", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering:", err)
	}
}

package harbas

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func (h *Harbas) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	if h.Debug {
		mux.Use(middleware.Logger)
	}
	mux.Use(middleware.Recoverer)

	return mux
}

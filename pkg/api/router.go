package api

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func NewRouter(api *roomTubeAPI) http.Handler {
	router := chi.NewRouter()

	// without
	router.Route("/", func(r chi.Router) {
		r.Post("/reg", api.Register)
		r.Post("/auth", api.Authorize)
	})

	// with middleware
	router.Route("/", func(r chi.Router) {
		r.Use(api.middleware.AuthMiddleware)
		r.Post("/user", api.GetUser)
	})

	return router
}

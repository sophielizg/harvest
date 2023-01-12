package routes

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/sophielizg/harvest/api/pkg/app"
)

var (
	currentApp app.App
)

func Init() (*chi.Mux, error) {
	// Initialize app
	var err error
	currentApp, err = app.Init()
	if err != nil {
		return nil, err
	}

	// Create router
	router := chi.NewRouter()

	// Add middlewares for all routes
	router.Use(
		middleware.Logger,
		middleware.RequestID,
		middleware.RedirectSlashes,
		middleware.Recoverer,
		render.SetContentType(render.ContentTypeJSON),
	)

	// Mount each route
	router.Route("/api/v1", func(r chi.Router) {
		r.Mount("/hello", HelloWorld())
	})

	return router, nil
}

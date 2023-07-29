package harvest

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/sophielizg/harvest/pkg/utils"
	"github.com/sophielizg/harvest/worker/routes"
)

func Routes() *chi.Mux {
	// Initialize app state
	app := &utils.App{}

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
		r.Mount("/hello", routes.HelloWorld(*app))
	})

	return router
}

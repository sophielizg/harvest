package routes

import (
	"fmt"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	_ "github.com/sophielizg/harvest/api/docs"
	"github.com/sophielizg/harvest/api/harvest"
	httpSwagger "github.com/swaggo/http-swagger"
)

var app *harvest.App

func CreateRouter(newApp *harvest.App, port string) (*chi.Mux, error) {
	app = newApp
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
		r.Mount("/crawls", CrawlRoutes())
		// Add parser types route
	})

	// Create swagger UI
	router.Get("/api/doc/*", httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprint("http://localhost", port, "/api/doc/doc.json")),
	))

	return router, nil
}

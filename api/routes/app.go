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

type App struct {
	CrawlService harvest.CrawlService
}

func (app *App) Router(port string) (*chi.Mux, error) {
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
		r.Mount("/crawls", app.CrawlRouter())
		// Add parser types route
	})

	// Create swagger UI
	router.Get("/api/doc/*", httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprint("http://localhost", port, "/api/doc/doc.json")),
	))

	return router, nil
}

package routes

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func (app *App) ParserRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/all", WriteErrorResponse(app.getParsers))
	// router.Post("/{crawlId}/parsers/add", ...)
	// router.Delete("/{crawlId}/parsers/{parserId}/delete", ...)
	// router.Post("/{crawlId}/parsers/{parserId}/tags/add", ...)
	// router.Delete("/{crawlId}/parsers/{parserId}/tags/delete", ...)

	return router
}

// getParsers godoc
// @Summary Get parsers endpoint
// @Description Get parsers for a crawl using its crawlId
// @Accept  json
// @Produce  json
// @Param crawlId path string true "Id of crawl"
// @Success 200 {object} []harvest.Parser
// @Failure 400 {object} ErrorResponse
// @Router /crawls/{crawlId}/parsers/all [get]
func (app *App) getParsers(r *http.Request) (interface{}, error) {
	crawlIdStr := chi.URLParam(r, "crawlId")
	crawlId, err := strconv.Atoi(crawlIdStr)
	if err != nil {
		return nil, err
	}
	return app.ParserService.Parsers(crawlId)
}

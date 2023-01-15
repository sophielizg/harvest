package routes

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func CrawlRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{crawlId}", WriteErrorResponse(GetCrawlById))
	// router.Get("/all", ...)
	// router.Get("/name/{name}", HelloWorldGet)
	// router.Post("/add", ...)

	// router.Delete("/{crawlId}/delete", ...)
	// router.Patch("/{crawlId}/start", ...)
	// router.Patch("/{crawlId}/stop", ...)
	// router.Patch("/{crawlId}/pause", ...)
	// router.Patch("/{crawlId}/unpause", ...)
	// router.Post("/{crawlId}/runners/add", ...)

	// router.Get("/{crawlId}/status", ...)
	// router.Get("/{crawlId}/results", ...)
	// router.Get("/{crawlId}/errors", ...)
	// router.Get("/{crawlId}/runs/{crawlRunId}/status", ...)
	// router.Get("/{crawlId}/runs/{crawlRunId}/results", ...)
	// router.Get("/{crawlId}/runs/{crawlRunId}/errors", ...)

	// router.Get("/{crawlId}/runs/all", ...)
	// router.Delete("/{crawlId}/runs/{crawlRunId}/delete", ...)

	// router.Get("/{crawlId}/requests/all", ...)
	// router.Post("/{crawlId}/requests/add", ...)
	// router.Delete("/{crawlId}/requests/{requestQueueId}/delete", ...)

	// router.Post("/{crawlId}/parsers/add", ...)
	// router.Get("/{crawlId}/parsers/all", ...)
	// router.Delete("/{crawlId}/parsers/{parserId}/delete", ...)
	// router.Post("/{crawlId}/parsers/{parserId}/tags/add", ...)
	// router.Delete("/{crawlId}/parsers/{parserId}/tags/delete", ...)

	return router
}

// GetCrawlById godoc
// @Summary Get crawl endpoint
// @Description Get crawl details using its crawlId
// @Accept  json
// @Produce  json
// @Param crawlId path string true "Id of crawl"
// @Success 200 {object} harvest.Crawl
// @Failure 400 {object} ErrorResponse
// @Router /crawls/{crawlId} [get]
func GetCrawlById(r *http.Request) (interface{}, error) {
	crawlIdStr := chi.URLParam(r, "crawlId")
	crawlId, err := strconv.Atoi(crawlIdStr)
	if err != nil {
		return nil, err
	}
	return app.CrawlService.Crawl(crawlId)
}

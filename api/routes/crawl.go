package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/sophielizg/harvest/api/harvest"
)

type AddCrawlResponse struct {
	CrawlId int `json:"crawlId"`
}

func (app *App) CrawlRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{crawlId}", WriteErrorResponse(app.getCrawlById))
	router.Get("/name/{name}", WriteErrorResponse(app.getCrawlByName))
	router.Get("/all", WriteErrorResponse(app.getCrawls))
	router.Post("/add", WriteErrorResponse(app.addCrawl))

	router.Post("/{crawlId}/update", WriteErrorResponse(app.updateCrawl))
	router.Delete("/{crawlId}/delete", WriteErrorResponse(app.deleteCrawl))
	router.Patch("/{crawlId}/start", WriteErrorResponse(app.startCrawl))
	router.Patch("/{crawlId}/stop", WriteErrorResponse(app.stopCrawl))
	router.Patch("/{crawlId}/pause", WriteErrorResponse(app.pauseCrawl))
	router.Patch("/{crawlId}/unpause", WriteErrorResponse(app.unpauseCrawl))
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

// getCrawlById godoc
// @Summary Get crawl endpoint
// @Description Get crawl details using its crawlId
// @Accept  json
// @Produce  json
// @Param crawlId path string true "Id of crawl"
// @Success 200 {object} harvest.Crawl
// @Failure 400 {object} ErrorResponse
// @Router /crawls/{crawlId} [get]
func (app *App) getCrawlById(r *http.Request) (interface{}, error) {
	crawlIdStr := chi.URLParam(r, "crawlId")
	crawlId, err := strconv.Atoi(crawlIdStr)
	if err != nil {
		return nil, err
	}
	return app.CrawlService.Crawl(crawlId)
}

// getCrawlById godoc
// @Summary Get crawl endpoint
// @Description Get crawl details using its name
// @Accept  json
// @Produce  json
// @Param name path string true "Name of crawl"
// @Success 200 {object} harvest.Crawl
// @Failure 400 {object} ErrorResponse
// @Router /crawls/name/{name} [get]
func (app *App) getCrawlByName(r *http.Request) (interface{}, error) {
	name := chi.URLParam(r, "name")
	return app.CrawlService.CrawlByName(name)
}

// getCrawls godoc
// @Summary Get crawl endpoint
// @Description Get details of all crawls
// @Accept  json
// @Produce  json
// @Success 200 {object} []harvest.Crawl
// @Failure 400 {object} ErrorResponse
// @Router /crawls/all [get]
func (app *App) getCrawls(r *http.Request) (interface{}, error) {
	return app.CrawlService.Crawls()
}

// addCrawl godoc
// @Summary Add crawl endpoint
// @Description Add a new crawl
// @Accept  json
// @Produce  json
// @Param request body harvest.CrawlFields true "Fields for crawl"
// @Success 200 {object} AddCrawlResponse
// @Failure 400 {object} ErrorResponse
// @Router /crawls/add [post]
func (app *App) addCrawl(r *http.Request) (interface{}, error) {
	var crawl harvest.CrawlFields
	err := json.NewDecoder(r.Body).Decode(&crawl)
	if err != nil {
		return nil, err
	}

	crawlId, err := app.CrawlService.AddCrawl(crawl)
	if err != nil {
		return nil, err
	}

	return AddCrawlResponse{CrawlId: crawlId}, nil
}

// updateCrawl godoc
// @Summary Update crawl endpoint
// @Description Update an existing crawl
// @Accept  json
// @Produce  json
// @Param crawlId path string true "Id of crawl"
// @Param request body harvest.CrawlFields true "Fields for crawl"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Router /crawls/{crawlId}/update [post]
func (app *App) updateCrawl(r *http.Request) (interface{}, error) {
	crawlIdStr := chi.URLParam(r, "crawlId")
	crawlId, err := strconv.Atoi(crawlIdStr)
	if err != nil {
		return nil, err
	}

	var crawl harvest.CrawlFields
	err = json.NewDecoder(r.Body).Decode(&crawl)
	if err != nil {
		return nil, err
	}

	err = app.CrawlService.UpdateCrawl(crawlId, crawl)
	if err != nil {
		return nil, err
	}

	return SuccessResponse{Message: "Successfully updated"}, nil
}

// deleteCrawl godoc
// @Summary Delete crawl endpoint
// @Description Delete crawl by its crawlId
// @Accept  json
// @Produce  json
// @Param crawlId path string true "Id of crawl"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Router /crawls/{crawlId}/delete [delete]
func (app *App) deleteCrawl(r *http.Request) (interface{}, error) {
	crawlIdStr := chi.URLParam(r, "crawlId")
	crawlId, err := strconv.Atoi(crawlIdStr)
	if err != nil {
		return nil, err
	}

	err = app.CrawlService.DeleteCrawl(crawlId)
	if err != nil {
		return nil, err
	}

	return SuccessResponse{Message: "Successfully deleted"}, nil
}

// stopCrawl godoc
// @Summary Stop crawl endpoint
// @Description Stop crawl by its crawlId
// @Accept  json
// @Produce  json
// @Param crawlId path string true "Id of crawl"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Router /crawls/{crawlId}/stop [patch]
func (app *App) stopCrawl(r *http.Request) (interface{}, error) {
	crawlIdStr := chi.URLParam(r, "crawlId")
	crawlId, err := strconv.Atoi(crawlIdStr)
	if err != nil {
		return nil, err
	}

	err = app.CrawlService.StopCrawl(crawlId)
	if err != nil {
		return nil, err
	}

	return SuccessResponse{Message: "Successfully stopped"}, nil
}

// startCrawl godoc
// @Summary Start crawl endpoint
// @Description Start crawl by its crawlId
// @Accept  json
// @Produce  json
// @Param crawlId path string true "Id of crawl"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Router /crawls/{crawlId}/start [patch]
func (app *App) startCrawl(r *http.Request) (interface{}, error) {
	crawlIdStr := chi.URLParam(r, "crawlId")
	crawlId, err := strconv.Atoi(crawlIdStr)
	if err != nil {
		return nil, err
	}

	err = app.CrawlService.StartCrawl(crawlId)
	if err != nil {
		return nil, err
	}

	return SuccessResponse{Message: "Successfully started"}, nil
}

// pauseCrawl godoc
// @Summary Pause crawl endpoint
// @Description Pause crawl by its crawlId
// @Accept  json
// @Produce  json
// @Param crawlId path string true "Id of crawl"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Router /crawls/{crawlId}/pause [patch]
func (app *App) pauseCrawl(r *http.Request) (interface{}, error) {
	crawlIdStr := chi.URLParam(r, "crawlId")
	crawlId, err := strconv.Atoi(crawlIdStr)
	if err != nil {
		return nil, err
	}

	err = app.CrawlService.PauseCrawl(crawlId)
	if err != nil {
		return nil, err
	}

	return SuccessResponse{Message: "Successfully paused"}, nil
}

// unpauseCrawl godoc
// @Summary Unpause crawl endpoint
// @Description Unpause crawl by its crawlId
// @Accept  json
// @Produce  json
// @Param crawlId path string true "Id of crawl"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Router /crawls/{crawlId}/unpause [patch]
func (app *App) unpauseCrawl(r *http.Request) (interface{}, error) {
	crawlIdStr := chi.URLParam(r, "crawlId")
	crawlId, err := strconv.Atoi(crawlIdStr)
	if err != nil {
		return nil, err
	}

	err = app.CrawlService.UnpauseCrawl(crawlId)
	if err != nil {
		return nil, err
	}

	return SuccessResponse{Message: "Successfully unpaused"}, nil
}

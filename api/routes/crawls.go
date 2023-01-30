package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/sophielizg/harvest/common/harvest"
)

type AddScraperResponse struct {
	ScraperId int `json:"scraperId"`
}

func (app *App) ScraperRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{scraperId}", WriteErrorResponse(app.getScraperById))
	router.Get("/name/{name}", WriteErrorResponse(app.getScraperByName))
	router.Get("/all", WriteErrorResponse(app.getScrapers))
	router.Post("/add", WriteErrorResponse(app.addScraper))

	router.Post("/{scraperId}/update", WriteErrorResponse(app.updateScraper))
	router.Delete("/{scraperId}/delete", WriteErrorResponse(app.deleteScraper))
	router.Patch("/{scraperId}/start", WriteErrorResponse(app.startScraper))
	router.Patch("/{scraperId}/stop", WriteErrorResponse(app.stopScraper))
	router.Patch("/{scraperId}/pause", WriteErrorResponse(app.pauseScraper))
	router.Patch("/{scraperId}/unpause", WriteErrorResponse(app.unpauseScraper))
	// router.Post("/{scraperId}/runners/add", ...)

	// router.Get("/{scraperId}/status", ...)
	// router.Get("/{scraperId}/results", ...)
	// router.Get("/{scraperId}/errors", ...)
	// router.Get("/{scraperId}/runs/{runId}/status", ...)
	// router.Get("/{scraperId}/runs/{runId}/results", ...)
	// router.Get("/{scraperId}/runs/{runId}/errors", ...)

	// router.Get("/{scraperId}/runs/all", ...)
	// router.Delete("/{scraperId}/runs/{runId}/delete", ...)

	// router.Get("/{scraperId}/requests/all", ...)
	// router.Post("/{scraperId}/requests/add", ...)
	// router.Delete("/{scraperId}/requests/{requestQueueId}/delete", ...)

	// router.Post("/{scraperId}/parsers/add", ...)
	// router.Get("/{scraperId}/parsers/all", ...)
	// router.Delete("/{scraperId}/parsers/{parserId}/delete", ...)
	// router.Post("/{scraperId}/parsers/{parserId}/tags/add", ...)
	// router.Delete("/{scraperId}/parsers/{parserId}/tags/delete", ...)

	return router
}

// getScraperById godoc
// @Summary Get crawl endpoint
// @Description Get crawl details using its scraperId
// @Tags crawls
// @Accept  json
// @Produce  json
// @Param scraperId path string true "Id of crawl"
// @Success 200 {object} harvest.Scraper
// @Failure 400 {object} ErrorResponse
// @Router /crawls/{scraperId} [get]
func (app *App) getScraperById(r *http.Request) (interface{}, error) {
	scraperIdStr := chi.URLParam(r, "scraperId")
	scraperId, err := strconv.Atoi(scraperIdStr)
	if err != nil {
		return nil, err
	}
	return app.ScraperService.Scraper(scraperId)
}

// getScraperById godoc
// @Summary Get crawl endpoint
// @Description Get crawl details using its name
// @Tags crawls
// @Accept  json
// @Produce  json
// @Param name path string true "Name of crawl"
// @Success 200 {object} harvest.Scraper
// @Failure 400 {object} ErrorResponse
// @Router /crawls/name/{name} [get]
func (app *App) getScraperByName(r *http.Request) (interface{}, error) {
	name := chi.URLParam(r, "name")
	return app.ScraperService.ScraperByName(name)
}

// getScrapers godoc
// @Summary Get crawl endpoint
// @Description Get details of all crawls
// @Tags crawls
// @Accept  json
// @Produce  json
// @Success 200 {object} []harvest.Scraper
// @Failure 400 {object} ErrorResponse
// @Router /crawls/all [get]
func (app *App) getScrapers(r *http.Request) (interface{}, error) {
	return app.ScraperService.Scrapers()
}

// addScraper godoc
// @Summary Add crawl endpoint
// @Description Add a new crawl
// @Tags crawls
// @Accept  json
// @Produce  json
// @Param request body harvest.ScraperFields true "Fields for crawl"
// @Success 200 {object} AddScraperResponse
// @Failure 400 {object} ErrorResponse
// @Router /crawls/add [post]
func (app *App) addScraper(r *http.Request) (interface{}, error) {
	var crawl harvest.ScraperFields
	err := json.NewDecoder(r.Body).Decode(&crawl)
	if err != nil {
		return nil, err
	}

	scraperId, err := app.ScraperService.AddScraper(crawl)
	if err != nil {
		return nil, err
	}

	return AddScraperResponse{ScraperId: scraperId}, nil
}

// updateScraper godoc
// @Summary Update crawl endpoint
// @Description Update an existing crawl
// @Tags crawls
// @Accept  json
// @Produce  json
// @Param scraperId path string true "Id of crawl"
// @Param request body harvest.ScraperFields true "Fields for crawl"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Router /crawls/{scraperId}/update [post]
func (app *App) updateScraper(r *http.Request) (interface{}, error) {
	scraperIdStr := chi.URLParam(r, "scraperId")
	scraperId, err := strconv.Atoi(scraperIdStr)
	if err != nil {
		return nil, err
	}

	var crawl harvest.ScraperFields
	err = json.NewDecoder(r.Body).Decode(&crawl)
	if err != nil {
		return nil, err
	}

	err = app.ScraperService.UpdateScraper(scraperId, crawl)
	if err != nil {
		return nil, err
	}

	return SuccessResponse{Message: "Successfully updated"}, nil
}

// deleteScraper godoc
// @Summary Delete crawl endpoint
// @Description Delete crawl by its scraperId
// @Tags crawls
// @Accept  json
// @Produce  json
// @Param scraperId path string true "Id of crawl"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Router /crawls/{scraperId}/delete [delete]
func (app *App) deleteScraper(r *http.Request) (interface{}, error) {
	scraperIdStr := chi.URLParam(r, "scraperId")
	scraperId, err := strconv.Atoi(scraperIdStr)
	if err != nil {
		return nil, err
	}

	err = app.ScraperService.DeleteScraper(scraperId)
	if err != nil {
		return nil, err
	}

	return SuccessResponse{Message: "Successfully deleted"}, nil
}

// stopScraper godoc
// @Summary Stop crawl endpoint
// @Description Stop crawl by its scraperId
// @Tags crawls
// @Accept  json
// @Produce  json
// @Param scraperId path string true "Id of crawl"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Router /crawls/{scraperId}/stop [patch]
func (app *App) stopScraper(r *http.Request) (interface{}, error) {
	scraperIdStr := chi.URLParam(r, "scraperId")
	scraperId, err := strconv.Atoi(scraperIdStr)
	if err != nil {
		return nil, err
	}

	err = app.ScraperService.StopScraper(scraperId)
	if err != nil {
		return nil, err
	}

	return SuccessResponse{Message: "Successfully stopped"}, nil
}

// startScraper godoc
// @Summary Start crawl endpoint
// @Description Start crawl by its scraperId
// @Tags crawls
// @Accept  json
// @Produce  json
// @Param scraperId path string true "Id of crawl"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Router /crawls/{scraperId}/start [patch]
func (app *App) startScraper(r *http.Request) (interface{}, error) {
	scraperIdStr := chi.URLParam(r, "scraperId")
	scraperId, err := strconv.Atoi(scraperIdStr)
	if err != nil {
		return nil, err
	}

	err = app.ScraperService.StartScraper(scraperId)
	if err != nil {
		return nil, err
	}

	return SuccessResponse{Message: "Successfully started"}, nil
}

// pauseScraper godoc
// @Summary Pause crawl endpoint
// @Description Pause crawl by its scraperId
// @Tags crawls
// @Accept  json
// @Produce  json
// @Param scraperId path string true "Id of crawl"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Router /crawls/{scraperId}/pause [patch]
func (app *App) pauseScraper(r *http.Request) (interface{}, error) {
	scraperIdStr := chi.URLParam(r, "scraperId")
	scraperId, err := strconv.Atoi(scraperIdStr)
	if err != nil {
		return nil, err
	}

	err = app.ScraperService.PauseScraper(scraperId)
	if err != nil {
		return nil, err
	}

	return SuccessResponse{Message: "Successfully paused"}, nil
}

// unpauseScraper godoc
// @Summary Unpause crawl endpoint
// @Description Unpause crawl by its scraperId
// @Tags crawls
// @Accept  json
// @Produce  json
// @Param scraperId path string true "Id of crawl"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Router /crawls/{scraperId}/unpause [patch]
func (app *App) unpauseScraper(r *http.Request) (interface{}, error) {
	scraperIdStr := chi.URLParam(r, "scraperId")
	scraperId, err := strconv.Atoi(scraperIdStr)
	if err != nil {
		return nil, err
	}

	err = app.ScraperService.UnpauseScraper(scraperId)
	if err != nil {
		return nil, err
	}

	return SuccessResponse{Message: "Successfully unpaused"}, nil
}

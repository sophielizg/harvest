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
	router.Get("/all", WriteErrorResponse(app.getScrapers))
	router.Post("/add", WriteErrorResponse(app.addScraper))

	router.Post("/{scraperId}/update", WriteErrorResponse(app.updateScraper))
	router.Delete("/{scraperId}/delete", WriteErrorResponse(app.deleteScraper))
	router.Get("/{scraperId}/parsers/all", WriteErrorResponse(app.getParsers))
	router.Post("/{scraperId}/parsers/add", WriteErrorResponse(app.addParser))
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
// @Summary Get scraper endpoint
// @Description Get scraper details using its scraperId
// @Tags scrapers
// @Accept  json
// @Produce  json
// @Param scraperId path string true "Id of scraper"
// @Success 200 {object} harvest.Scraper
// @Failure 400 {object} ErrorResponse
// @Router /scrapers/{scraperId} [get]
func (app *App) getScraperById(r *http.Request) (interface{}, error) {
	scraperIdStr := chi.URLParam(r, "scraperId")
	scraperId, err := strconv.Atoi(scraperIdStr)
	if err != nil {
		return nil, err
	}
	return app.ScraperService.Scraper(scraperId)
}

// getScrapers godoc
// @Summary Get scraper endpoint
// @Description Get details of all scrapers
// @Tags scrapers
// @Accept  json
// @Produce  json
// @Success 200 {object} []harvest.Scraper
// @Failure 400 {object} ErrorResponse
// @Router /scrapers/all [get]
func (app *App) getScrapers(r *http.Request) (interface{}, error) {
	return app.ScraperService.Scrapers()
}

// addScraper godoc
// @Summary Add scraper endpoint
// @Description Add a new scraper
// @Tags scrapers
// @Accept  json
// @Produce  json
// @Param request body harvest.ScraperFields true "Fields for scraper"
// @Success 200 {object} AddScraperResponse
// @Failure 400 {object} ErrorResponse
// @Router /scrapers/add [post]
func (app *App) addScraper(r *http.Request) (interface{}, error) {
	var scraper harvest.ScraperFields
	err := json.NewDecoder(r.Body).Decode(&scraper)
	if err != nil {
		return nil, err
	}

	scraperId, err := app.ScraperService.AddScraper(scraper)
	if err != nil {
		return nil, err
	}

	return AddScraperResponse{ScraperId: scraperId}, nil
}

// updateScraper godoc
// @Summary Update scraper endpoint
// @Description Update an existing scraper
// @Tags scrapers
// @Accept  json
// @Produce  json
// @Param scraperId path string true "Id of scraper"
// @Param request body harvest.ScraperFields true "Fields for scraper"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Router /scrapers/{scraperId}/update [post]
func (app *App) updateScraper(r *http.Request) (interface{}, error) {
	scraperIdStr := chi.URLParam(r, "scraperId")
	scraperId, err := strconv.Atoi(scraperIdStr)
	if err != nil {
		return nil, err
	}

	var scraper harvest.ScraperFields
	err = json.NewDecoder(r.Body).Decode(&scraper)
	if err != nil {
		return nil, err
	}

	err = app.ScraperService.UpdateScraper(scraperId, scraper)
	if err != nil {
		return nil, err
	}

	return SuccessResponse{Message: "Successfully updated"}, nil
}

// deleteScraper godoc
// @Summary Delete scraper endpoint
// @Description Delete scraper by its scraperId
// @Tags scrapers
// @Accept  json
// @Produce  json
// @Param scraperId path string true "Id of scraper"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Router /scrapers/{scraperId}/delete [delete]
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

// getParsers godoc
// @Summary Get parsers endpoint
// @Description Get parsers for a scraper using its scraperId
// @Tags parsers
// @Accept  json
// @Produce  json
// @Param scraperId path string true "Id of scraper"
// @Success 200 {object} []harvest.Parser
// @Failure 400 {object} ErrorResponse
// @Router /scrapers/{scraperId}/parsers/all [get]
func (app *App) getParsers(r *http.Request) (interface{}, error) {
	scraperIdStr := chi.URLParam(r, "scraperId")
	scraperId, err := strconv.Atoi(scraperIdStr)
	if err != nil {
		return nil, err
	}
	return app.ParserService.Parsers(scraperId)
}

// addParser godoc
// @Summary Add parser endpoint
// @Description Add parser to a scraper
// @Tags parsers
// @Accept  json
// @Produce  json
// @Param scraperId path string true "Id of scraper"
// @Param request body harvest.ParserFields true "Fields for parser"
// @Success 200 {object} AddParserResponse
// @Failure 400 {object} ErrorResponse
// @Router /scrapers/{scraperId}/parsers/add [post]
func (app *App) addParser(r *http.Request) (interface{}, error) {
	scraperIdStr := chi.URLParam(r, "scraperId")
	scraperId, err := strconv.Atoi(scraperIdStr)
	if err != nil {
		return nil, err
	}

	var parser harvest.ParserFields
	err = json.NewDecoder(r.Body).Decode(&parser)
	if err != nil {
		return nil, err
	}

	parserId, err := app.ParserService.AddParser(scraperId, parser)
	if err != nil {
		return nil, err
	}

	return AddParserResponse{ParserId: parserId}, nil
}

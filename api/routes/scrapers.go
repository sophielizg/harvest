package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/sophielizg/harvest/common/harvest"
)

// Request bodies
type AddRequestBody struct {
	Url    string      `json:"url"`
	Method string      `json:"method"`
	Body   interface{} `json:"body"`
}

// Responses
type AddScraperResponse struct {
	ScraperId int `json:"scraperId"`
}

type AddParserResponse struct {
	ParserId int `json:"parserId"`
}

type CreateRunResponse struct {
	RunId int `json:"runId"`
}

type EnqueueRunnerResponse struct {
	RunnerId int `json:"runnerId"`
}

type EnqueueRequestResponse struct {
	RequestQueueId int `json:"requestQueueId"`
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

	router.Post("/{scraperId}/runs/create", WriteErrorResponse(app.createRun))
	router.Post("/{scraperId}/runners/current/enqueue",
		WriteErrorResponse(app.enqueueRunnerCurrentRun))
	// router.Post("/{scraperId}/runners/add", ...)

	router.Post("/{scraperId}/requests/queue/start/add",
		WriteErrorResponse(app.addStartingRequest))

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

// createRun godoc
// @Summary Create run endpoint
// @Description Add run to a scraper
// @Tags runs
// @Accept  json
// @Produce  json
// @Param scraperId path string true "Id of scraper"
// @Success 200 {object} CreateRunResponse
// @Failure 400 {object} ErrorResponse
// @Router /scrapers/{scraperId}/runs/create [post]
func (app *App) createRun(r *http.Request) (interface{}, error) {
	scraperIdStr := chi.URLParam(r, "scraperId")
	scraperId, err := strconv.Atoi(scraperIdStr)
	if err != nil {
		return nil, err
	}

	runId, err := app.RunService.CreateRun(scraperId)
	if err != nil {
		return nil, err
	}

	return CreateRunResponse{RunId: runId}, nil
}

// enqueueRunnerCurrentRun godoc
// @Summary Enqueue runner to current run endpoint
// @Description Enqueue a runner to the current run of a scraper
// @Tags runners
// @Accept  json
// @Produce  json
// @Param scraperId path string true "Id of scraper"
// @Success 200 {object} EnqueueRunnerResponse
// @Failure 400 {object} ErrorResponse
// @Router /scrapers/{scraperId}/runners/current/enqueue [post]
func (app *App) enqueueRunnerCurrentRun(r *http.Request) (interface{}, error) {
	scraperIdStr := chi.URLParam(r, "scraperId")
	scraperId, err := strconv.Atoi(scraperIdStr)
	if err != nil {
		return nil, err
	}

	runnerId, err := app.RunnerQueueService.EnqueueRunnerForCurrentRun(scraperId)
	if err != nil {
		return nil, err
	}

	return EnqueueRunnerResponse{RunnerId: runnerId}, nil
}

// addStartingRequest godoc
// @Summary Add starting request to scraper
// @Description Add a new starting request to a scraper
// @Tags requestqueue
// @Accept  json
// @Produce  json
// @Param scraperId path string true "Id of scraper"
// @Param request body AddRequestBody true "Request to send"
// @Success 200 {object} EnqueueRequestResponse
// @Failure 400 {object} ErrorResponse
// @Router /scrapers/{scraperId}/requests/queue/start/add [post]
func (app *App) addStartingRequest(r *http.Request) (interface{}, error) {
	scraperIdStr := chi.URLParam(r, "scraperId")
	scraperId, err := strconv.Atoi(scraperIdStr)
	if err != nil {
		return nil, err
	}

	var requestToSend AddRequestBody
	err = json.NewDecoder(r.Body).Decode(&requestToSend)
	if err != nil {
		return nil, err
	}

	requestBlob, err := json.Marshal(requestToSend)
	if err != nil {
		return nil, err
	}

	requestQueueId, err := app.RequestQueueService.AddStartingRequest(scraperId, requestBlob)
	if err != nil {
		return nil, err
	}

	return EnqueueRequestResponse{RequestQueueId: requestQueueId}, nil
}

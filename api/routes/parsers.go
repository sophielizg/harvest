package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/sophielizg/harvest/common/harvest"
)

type AddParserResponse struct {
	ParserId int `json:"parserId"`
}

func (app *App) ParserRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/all", WriteErrorResponse(app.getParsers))
	router.Post("/add", WriteErrorResponse(app.addParser))
	router.Delete("/{parserId}/delete", WriteErrorResponse(app.deleteParser))
	router.Post("/{parserId}/tags/add/{name}", WriteErrorResponse(app.addParserTag))
	router.Delete("/{parserId}/tags/delete/{name}", WriteErrorResponse(app.deleteParserTag))

	return router
}

// getParsers godoc
// @Summary Get parsers endpoint
// @Description Get parsers for a crawl using its crawlId
// @Tags parsers
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

// addParser godoc
// @Summary Add parser endpoint
// @Description Add parser to a crawl
// @Tags parsers
// @Accept  json
// @Produce  json
// @Param crawlId path string true "Id of crawl"
// @Param request body harvest.ParserFields true "Fields for parser"
// @Success 200 {object} AddParserResponse
// @Failure 400 {object} ErrorResponse
// @Router /crawls/{crawlId}/parsers/add [post]
func (app *App) addParser(r *http.Request) (interface{}, error) {
	crawlIdStr := chi.URLParam(r, "crawlId")
	crawlId, err := strconv.Atoi(crawlIdStr)
	if err != nil {
		return nil, err
	}

	var parser harvest.ParserFields
	err = json.NewDecoder(r.Body).Decode(&parser)
	if err != nil {
		return nil, err
	}

	parserId, err := app.ParserService.AddParser(crawlId, parser)
	if err != nil {
		return nil, err
	}

	return AddParserResponse{ParserId: parserId}, nil
}

// deleteParser godoc
// @Summary Delete parser endpoint
// @Description Delete parser by its parserId
// @Tags parsers
// @Accept  json
// @Produce  json
// @Param crawlId path string true "Id of crawl"
// @Param parserId path string true "Id of parser"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Router /crawls/{crawlId}/parsers/{parserId}/delete [delete]
func (app *App) deleteParser(r *http.Request) (interface{}, error) {
	parserIdStr := chi.URLParam(r, "parserId")
	parserId, err := strconv.Atoi(parserIdStr)
	if err != nil {
		return nil, err
	}

	err = app.ParserService.DeleteParser(parserId)
	if err != nil {
		return nil, err
	}

	return SuccessResponse{Message: "Successfully deleted"}, nil
}

// addParserTag godoc
// @Summary Add parser tag endpoint
// @Description Add a tag to a parser
// @Tags parsers
// @Accept  json
// @Produce  json
// @Param crawlId path string true "Id of crawl"
// @Param parserId path string true "Id of parser"
// @Param name path string true "Name of tag"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Router /crawls/{crawlId}/parsers/{parserId}/tags/add/{name} [post]
func (app *App) addParserTag(r *http.Request) (interface{}, error) {
	tagName := chi.URLParam(r, "name")
	parserIdStr := chi.URLParam(r, "parserId")
	parserId, err := strconv.Atoi(parserIdStr)
	if err != nil {
		return nil, err
	}

	err = app.ParserService.AddParserTag(parserId, tagName)
	if err != nil {
		return nil, err
	}

	return SuccessResponse{Message: "Successfully added"}, nil
}

// deleteParserTag godoc
// @Summary Delete parser tag endpoint
// @Description Delete a tag from a parser
// @Tags parsers
// @Accept  json
// @Produce  json
// @Param crawlId path string true "Id of crawl"
// @Param parserId path string true "Id of parser"
// @Param name path string true "Name of tag"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Router /crawls/{crawlId}/parsers/{parserId}/tags/delete/{name} [post]
func (app *App) deleteParserTag(r *http.Request) (interface{}, error) {
	tagName := chi.URLParam(r, "name")
	parserIdStr := chi.URLParam(r, "parserId")
	parserId, err := strconv.Atoi(parserIdStr)
	if err != nil {
		return nil, err
	}

	err = app.ParserService.DeleteParserTag(parserId, tagName)
	if err != nil {
		return nil, err
	}

	return SuccessResponse{Message: "Successfully deleted"}, nil
}

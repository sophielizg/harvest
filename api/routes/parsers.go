package routes

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func (app *App) ParserRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Delete("/{parserId}/delete", WriteErrorResponse(app.deleteParser))
	router.Post("/{parserId}/tags/add/{name}", WriteErrorResponse(app.addParserTag))
	router.Delete("/{parserId}/tags/delete/{name}", WriteErrorResponse(app.deleteParserTag))

	return router
}

// deleteParser godoc
// @Summary Delete parser endpoint
// @Description Delete parser by its parserId
// @Tags parsers
// @Accept  json
// @Produce  json
// @Param parserId path string true "Id of parser"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Router /parsers/{parserId}/delete [delete]
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
// @Param parserId path string true "Id of parser"
// @Param name path string true "Name of tag"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Router /parsers/{parserId}/tags/add/{name} [post]
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
// @Param parserId path string true "Id of parser"
// @Param name path string true "Name of tag"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Router /parsers/{parserId}/tags/delete/{name} [post]
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

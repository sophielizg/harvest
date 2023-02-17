package routes

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (app *App) RunnerRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Post("/create", writeErrorResponse(app.createRunner))

	return router
}

// createRunner godoc
// @Summary Create runner endpoint
// @Description Create a new runner
// @Tags runners
// @Accept  json
// @Produce  json
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Router /runners/create [post]
func (app *App) createRunner(r *http.Request) (interface{}, error) {
	err := app.RunnerService.CreateNewRunner()
	if err != nil {
		return nil, err
	}

	return SuccessResponse{Message: "Successfully created"}, nil
}

package routes

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/sophielizg/harvest/api/pkg/helloworld"
)

func HelloWorld() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", HelloWorldGet)
	return router
}

// HelloWorldGet godoc
// @Summary Hello world endpoint
// @Description Greet someone by name or say goodbye
// @Accept  json
// @Produce  json
// @Param name query string false "Person to greet"
// @Success 200 {object} helloworld.Success
// @Failure 400 {object} helloworld.Failure
// @Router /hello [get]
func HelloWorldGet(w http.ResponseWriter, r *http.Request) {
	// Get params
	name := r.URL.Query().Get("name")

	// Call func with business logic
	response := helloworld.SayHello(currentApp, name)
	if _, isType := response.(helloworld.Success); !isType {
		render.Status(r, 400)
	}

	render.JSON(w, r, response)
}

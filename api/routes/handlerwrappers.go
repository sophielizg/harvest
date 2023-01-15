package routes

import (
	"net/http"

	"github.com/go-chi/render"
)

type HandlerFunc func(r *http.Request) (interface{}, error)

func WriteErrorResponse(handler HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response, err := handler(r)
		if err != nil {
			response = ErrorResponse{
				reason: err.Error(),
			}
			render.Status(r, 400)
		}

		render.JSON(w, r, response)
	})
}

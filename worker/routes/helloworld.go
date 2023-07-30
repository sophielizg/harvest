package routes

// import (
// 	"net/http"

// 	"github.com/go-chi/chi"
// 	"github.com/go-chi/render"
// 	"github.com/sophielizg/harvest/pkg/helloworld"
// 	"github.com/sophielizg/harvest/pkg/models"
// 	"github.com/sophielizg/harvest/pkg/utils"
// )

// func HelloWorld(app utils.App) *chi.Mux {
// 	router := chi.NewRouter()

// 	// Use the app object to wrap each handler in closure (if global state is needed)
// 	router.Get("/", app.WrapHandler(HelloWorldGet))

// 	return router
// }

// // HelloWorldGet godoc
// // @Summary Hello world endpoint
// // @Description Greet someone by name or say goodbye
// // @Accept  json
// // @Produce  json
// // @Param name query string false "Person to greet"
// // @Success 200 {object} models.HelloWorldSuccess
// // @Failure 400 {object} models.HelloWorldFailure
// // @Router /hello [get]
// func HelloWorldGet(app utils.App, w http.ResponseWriter, r *http.Request) {
// 	// Get params
// 	name := r.URL.Query().Get("name")

// 	// Call func with business logic
// 	response := helloworld.SayHello(app, name)
// 	if _, isType := response.(models.HelloWorldSuccess); !isType {
// 		render.Status(r, 400)
// 	}

// 	render.JSON(w, r, response)
// }

package app

import "net/http"

type App struct{} // Store global state like DB connections

type Handler func(app App, w http.ResponseWriter, r *http.Request)

func Init() (App, error) {
	return App{}, nil
}

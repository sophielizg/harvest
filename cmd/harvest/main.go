package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/sophielizg/harvest/docs"
	"github.com/sophielizg/harvest/routes"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title harvest
// @version 1.0
// @description Configureable web scraper to crawl and collect data from any website

// @BasePath /api/v1
func main() {
	// Grab PORT env variable
	port := fmt.Sprint(":", os.Getenv("PORT"))
	if port == ":" {
		port = ":8080"
	}

	// Initialize server
	router, err := routes.Init()
	if err != nil {
		log.Fatal(err)
	}

	// Create swagger UI
	router.Get("/api/doc/*", httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprint("http://localhost", port, "/api/doc/doc.json")),
	))

	// Start server
	log.Println(fmt.Sprint("Server running on http://localhost", port))
	log.Fatal(http.ListenAndServe(port, router))
}

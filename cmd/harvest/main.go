package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/sophielizg/harvest/docs"
	"github.com/sophielizg/harvest/harvest"
)

// @title harvest
// @version 1.0
// @description A framework for distributed web scraping

// @BasePath /api/v1
func main() {
	// Grab PORT env variable
	port := fmt.Sprint(":", os.Getenv("PORT"))
	if port == ":" {
		port = ":5000"
	}

	// Initialize server
	router := harvest.Routes()

	// Create swagger UI
	router.Get("/api/doc/*", httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprint("http://localhost", port, "/api/doc/doc.json")),
	))

	// Start server
	log.Println(fmt.Sprint("Server running on http://localhost", port))
	log.Fatal(http.ListenAndServe(port, router))
}

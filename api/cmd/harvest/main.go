package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sophielizg/harvest/api/config"
	"github.com/sophielizg/harvest/api/harvest"
	"github.com/sophielizg/harvest/api/mysql"
	"github.com/sophielizg/harvest/api/routes"
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

	// Create app
	app := harvest.App{
		configService: config.ConfigService{},
		crawlService:  mysql.CrawlService{},
	}

	// Connect db
	err := mysql.Open(app)
	if err != nil {
		log.Fatal(err)
	}
	defer mysql.Close()

	// Initialize server
	router, err := routes.CreateRouter(app, port)
	if err != nil {
		log.Fatal(err)
	}

	// Start server
	log.Println(fmt.Sprint("Server running on http://localhost", port))
	log.Fatal(http.ListenAndServe(port, router))
}

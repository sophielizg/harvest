package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sophielizg/harvest/api/config"
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

	// Create config service
	configService := &config.ConfigService{}

	// Connect to db
	db, err := mysql.OpenDb(configService)
	if err != nil {
		log.Fatal(err)
	}
	defer mysql.CloseDb(db)

	crawlService := &mysql.CrawlService{Db: db}

	// Initialize server
	app := routes.App{
		CrawlService: crawlService,
	}
	router, err := app.Router(port)
	if err != nil {
		log.Fatal(err)
	}

	// Start server
	log.Println(fmt.Sprint("Server running on http://localhost", port))
	log.Fatal(http.ListenAndServe(port, router))
}

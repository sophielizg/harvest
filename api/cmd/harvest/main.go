package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	// Import any external types used by swag
	_ "github.com/sophielizg/harvest/common/harvest"

	"github.com/sophielizg/harvest/api/routes"
	"github.com/sophielizg/harvest/common/config"
	"github.com/sophielizg/harvest/common/mysql"
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

	crawlService := &mysql.ScraperService{Db: db}
	parserService := &mysql.ParserService{Db: db}

	// Initialize server
	app := routes.App{
		ScraperService: crawlService,
		ParserService:  parserService,
	}
	router, err := app.Router(port)
	if err != nil {
		log.Fatal(err)
	}

	// Start server
	log.Println(fmt.Sprint("Server running on http://localhost", port))
	log.Fatal(http.ListenAndServe(port, router))
}

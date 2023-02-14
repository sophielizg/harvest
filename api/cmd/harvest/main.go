package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	// Import any external types used by swag
	_ "github.com/sophielizg/harvest/common/harvest"

	"github.com/sophielizg/harvest/api/routes"
	"github.com/sophielizg/harvest/common/local"
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

	// Create local services
	localServices, err := local.Init()
	if err != nil {
		log.Fatal(err)
	}

	// Create db connected services
	mysqlServices, err := mysql.Init(localServices.ConfigService)
	if err != nil {
		log.Fatal(err)
	}
	defer mysqlServices.Close()

	// Initialize server
	app := routes.App{
		RunnerService:       localServices.RunnerService,
		ScraperService:      mysqlServices.ScraperService,
		ParserService:       mysqlServices.ParserService,
		RunService:          mysqlServices.RunService,
		RunnerQueueService:  mysqlServices.RunnerQueueService,
		RequestQueueService: mysqlServices.RequestQueueService,
	}
	router, err := app.Router(port)
	if err != nil {
		log.Fatal(err)
	}

	// Start server
	log.Println(fmt.Sprint("Server running on http://localhost", port))
	log.Fatal(http.ListenAndServe(port, router))
}

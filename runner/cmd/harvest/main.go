package main

import (
	"log"

	"github.com/sophielizg/harvest/common/config"
	"github.com/sophielizg/harvest/common/mysql"
	"github.com/sophielizg/harvest/runner/colly"
)

func main() {
	// Create config service
	configService := &config.ConfigService{}

	// Connect to db
	db, err := mysql.OpenDb(configService)
	if err != nil {
		log.Fatal(err)
	}
	defer mysql.CloseDb(db)

	crawlService := &mysql.CrawlService{Db: db}
	scrapeService := &mysql.ScrapeService{Db: db}
	parserService := &mysql.ParserService{Db: db}
	resultService := &mysql.ResultService{Db: db}
	errorService := &mysql.ErrorService{Db: db}

	// Initialize runner
	app := colly.App{
		CrawlService:  crawlService,
		ScrapeService: scrapeService,
		ParserService: parserService,
		ResultService: resultService,
		ErrorService:  errorService,
	}

}

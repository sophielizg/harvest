package main

import (
	"log"

	"github.com/sophielizg/harvest/common/local"
	"github.com/sophielizg/harvest/common/mysql"
	"github.com/sophielizg/harvest/runner/colly"
	"github.com/sophielizg/harvest/runner/colly/parsers"
	"github.com/sophielizg/harvest/runner/colly/storage"
)

func main() {
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

	// Initialize runner
	runner := colly.Runner{
		ScraperService:     mysqlServices.ScraperService,
		RunnerQueueService: mysqlServices.RunnerQueueService,
		RequestService:     mysqlServices.RequestService,
		StorageServices: storage.StorageServices{
			CookieService:       mysqlServices.CookieService,
			VisitedService:      mysqlServices.VisitedService,
			RequestQueueService: mysqlServices.RequestQueueService,
		},
		ParsersServices: parsers.ParsersServices{
			ParserService: mysqlServices.ParserService,
			ResultService: mysqlServices.ResultService,
			ErrorService:  mysqlServices.ErrorService,
		},
	}

	err = runner.Run()
	if err != nil {
		log.Fatal(err)
	}
}

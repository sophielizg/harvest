package main

import (
	harvest "github.com/sophielizg/harvest/common"
	"github.com/sophielizg/harvest/common/local"
	"github.com/sophielizg/harvest/common/mysql"
	"github.com/sophielizg/harvest/common/zap"
	"github.com/sophielizg/harvest/runner/colly"
	"github.com/sophielizg/harvest/runner/colly/common"
	"github.com/sophielizg/harvest/runner/colly/parsers"
	"github.com/sophielizg/harvest/runner/colly/storage"
)

func main() {
	// Create logger
	logger := zap.Init()
	defer logger.Close()

	// Create local services
	localServices, err := local.Init(logger)
	if err != nil {
		logger.WithFields(harvest.LogFields{
			"error": err,
		}).Fatal("Could not create local services")
	}

	// Create db connected services
	mysqlServices, err := mysql.Init(localServices.ConfigService)
	if err != nil {
		logger.WithFields(harvest.LogFields{
			"error": err,
		}).Fatal("Could not create mysql services")
	}
	defer mysqlServices.Close()

	// Initialize runner
	runner := colly.Runner{
		SharedFields: common.SharedFields{
			Logger: logger,
		},
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
		logger.WithFields(harvest.LogFields{
			"error": err,
		}).Fatal("A fatal error ocurred while within the runner")
	}
}

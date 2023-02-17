package mysql

import (
	"database/sql"

	harvest "github.com/sophielizg/harvest/common"
)

type MysqlServices struct {
	db                  *sql.DB
	CookieService       *CookieService
	ErrorService        *ErrorService
	ParserService       *ParserService
	RequestService      *RequestService
	RequestQueueService *RequestQueueService
	ResultService       *ResultService
	RunnerQueueService  *RunnerQueueService
	ScraperService      *ScraperService
	VisitedService      *VisitedService
	RunService          *RunService
}

func (m *MysqlServices) Close() {
	CloseDb(m.db)
}

func Init(configService harvest.ConfigService) (*MysqlServices, error) {
	db, err := OpenDb(configService)
	if err != nil {
		return nil, err
	}

	return &MysqlServices{
		db,
		&CookieService{db},
		&ErrorService{db},
		&ParserService{db},
		&RequestService{db},
		&RequestQueueService{db},
		&ResultService{db},
		&RunnerQueueService{db},
		&ScraperService{db},
		&VisitedService{db},
		&RunService{db},
	}, nil
}

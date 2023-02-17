package common

import "github.com/sophielizg/harvest/common"

type SharedIds struct {
	ScraperId int
	RunId     int
	RunnerId  int
}

type SharedFields struct {
	SharedIds
	Logger common.Logger
}

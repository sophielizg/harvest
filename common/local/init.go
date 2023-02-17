package local

import harvest "github.com/sophielizg/harvest/common"

type LocalServices struct {
	ConfigService *ConfigService
	RunnerService *RunnerService
}

func Init(logger harvest.Logger) (*LocalServices, error) {
	configService := &ConfigService{}
	if err := configService.Init(); err != nil {
		return nil, err
	}

	return &LocalServices{
		configService,
		&RunnerService{
			logger,
		},
	}, nil
}

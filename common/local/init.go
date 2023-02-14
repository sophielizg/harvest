package local

type LocalServices struct {
	ConfigService *ConfigService
	RunnerService *RunnerService
}

func Init() (*LocalServices, error) {
	configService := &ConfigService{}
	if err := configService.Init(); err != nil {
		return nil, err
	}

	return &LocalServices{
		configService,
		&RunnerService{},
	}, nil
}

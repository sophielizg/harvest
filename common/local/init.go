package local

type LocalServices struct {
	ConfigService *ConfigService
	RunnerService *RunnerService
}

func Init() *LocalServices {
	return &LocalServices{
		ConfigService{},
		RunnerService{},
	}
}

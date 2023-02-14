package docker

type DockerServices struct {
	RunnerService *RunnerService
}

func Init() (*DockerServices, error) {
	return &DockerServices{
		&RunnerService{},
	}, nil
}

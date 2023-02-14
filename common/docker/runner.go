package docker

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
)

var (
	runnerImageName = "harvest-runner"
)

type RunnerService struct{}

func (r *RunnerService) CreateNewRunner() error {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}
	defer cli.Close()

	// TODO: pull image once it is hosted somewhere
	// out, err := cli.ImagePull(ctx, runnerImageName, types.ImagePullOptions{})
	// if err != nil {
	// 	return err
	// }
	// defer out.Close()

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: runnerImageName,
		Env:   []string{"ENV=docker"},
	}, nil, &network.NetworkingConfig{
		EndpointsConfig: map[string]*network.EndpointSettings{
			"mysql": {},
		},
	}, nil, "")
	if err != nil {
		return err
	}

	return cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
}

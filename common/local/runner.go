package local

import (
	"fmt"
	"os"
	"os/exec"
)

var (
	runnerDir = "../../runner"
)

type RunnerService struct{}

func (r *RunnerService) CreateNewRunner() error {
	cmd := exec.Command(runnerDir + "/runner")

	go func() {
		err := cmd.Run()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Run error: %s\n", err)
		}
	}()

	return nil
}

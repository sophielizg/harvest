package local

import (
	"os"
	"os/exec"

	"github.com/sophielizg/harvest/common"
)

var (
	runnerDir = "../runner"
)

type RunnerService struct {
	logger common.Logger
}

func (r *RunnerService) CreateNewRunner() error {
	cmd := exec.Command(runnerDir + "/runner")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	go func() {
		err := cmd.Run()

		if err != nil {
			r.logger.WithFields(common.LogFields{
				"error": err,
			}).Error("An error ocurred inside the runner")
		}
	}()

	return nil
}

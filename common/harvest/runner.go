package harvest

type RunnerService interface {
	CreateNewRunner(runId int) error
}

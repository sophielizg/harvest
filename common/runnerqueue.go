package common

type Runner struct {
	ScraperId int `json:"scraperId"`
	RunId     int `json:"runId"`
	RunnerId  int `json:"runnerId"`
}

type RunnerQueueService interface {
	EnqueueRunnerForRun(runId int) (int, error)
	EnqueueRunnerForCurrentRun(scraperId int) (int, error)
	DequeueRunner() (*Runner, error)
	EndRunner(runnerId int) error
}

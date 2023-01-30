package harvest

type Runner struct {
	ScraperId int `json:"scraperId"`
	RunId     int `json:"runId"`
	RunnerId  int `json:"runnerId"`
}

type RunnerQueueService interface {
	EnqueueRunner(runId int) (int, error)
	DequeueRunner() (*Runner, error)
	EndRunner(runnerId int) error
}

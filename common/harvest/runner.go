package harvest

type RunnerService interface {
	AddRunnerToCrawl(crawlId int) error
}

package harvest

type VisitedService interface {
	GetIsVisited(runId int, requestHash uint64) (bool, error)
	SetIsVisited(runId int, requestHash uint64) error
}

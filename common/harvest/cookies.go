package harvest

type CookieService interface {
	SetCookies(runId int, host string, value string) error
	GetCookies(runId int, host string) (string, error)
}

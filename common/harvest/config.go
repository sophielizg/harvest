package harvest

type ConfigService interface {
	Value(key string) (string, error)
}

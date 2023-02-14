package harvest

type ConfigService interface {
	Value(keys ...string) ([]byte, error)
}

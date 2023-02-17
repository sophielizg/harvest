package common

type ConfigService interface {
	Value(keys ...string) ([]byte, error)
}

package local

type ConfigService struct{}

func (c *ConfigService) Value(key string) (string, error) {
	return `{
		"user": "harvest",
		"password": "changeme",
		"protocol": "tcp",
		"host": "localhost",
		"port": 3306,
		"dbname": "harvest"
	}`, nil
}

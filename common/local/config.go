package local

import (
	"encoding/json"
	"os"

	"github.com/sophielizg/harvest/common/utils"
)

var (
	configDir = "./config"
)

type ConfigService struct {
	config interface{}
}

func (c *ConfigService) Init() error {
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}

	configBytes, err := os.ReadFile(configDir + "/" + env + ".json")
	if err != nil {
		return err
	}

	return json.Unmarshal(configBytes, &c.config)
}

func (c *ConfigService) Value(keys ...string) ([]byte, error) {
	ptr, err := utils.PointerFromJson(&c.config, keys)
	if err != nil {
		return nil, err
	}

	return json.Marshal(*ptr)
}

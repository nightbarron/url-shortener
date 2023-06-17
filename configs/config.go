package configs

import (
	"encoding/json"
	_ "gin_template/models"
	"io/ioutil"
	"os"
)

type Config struct {
	Database struct {
		User     string `json:"user"`
		Password string `json:"password"`
		Database string `json:"database"`
	} `json:"database"`
}

func (c *Config) Load(configFile string) error {
	file, err := os.Open(configFile)

	if err != nil {
		return err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, c)
	if err != nil {
		return err
	}

	return nil
}

package configs

import (
	"encoding/json"
	"io/ioutil"
	"os"
	_ "url-shortener/models"
)

type GlobalConfig struct {
	MySQL struct {
		Host         string `json:"HOST"`
		Port         string `json:"PORT"`
		User         string `json:"USER"`
		Pass         string `json:"PASS"`
		DBNamePrefix string `json:"DB_NAME_PREFIX"`
		NumberOfDB   int    `json:"NUMBER_OF_DB"`
	} `json:"MYSQL"`
	Redis struct {
		Host string `json:"HOST"`
		Port string `json:"PORT"`
		Pass string `json:"PASS"`
	} `json:"REDIS"`
	ShortUrlDomain string `json:"SHORT_URL_DOMAIN"`
}

func (c *GlobalConfig) Load(configFile string) error {
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

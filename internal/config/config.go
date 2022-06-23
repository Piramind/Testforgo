package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Port      string `yaml:"port"`
	SecretKey string `yaml:"secretKey"`
	MongoUrl  string `yaml:"mongoUrl"`
}

func NewConfig(path string) (*Config, error) {
	c := new(Config)
	yamlFile, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(yamlFile, c)

	if err != nil {
		return nil, err
	}

	return c, nil
}

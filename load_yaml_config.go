package main

import (
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

func LoadYamlConfig(configpath string) (*Config, error) {

	f, err := os.Open(configpath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var result *Config

	content, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

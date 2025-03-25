package main

import (
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

func LoadYamlConfig(configpath string) (string, error) {

	f, err := os.Open(configpath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	tmp := map[string]interface{}{}

	content, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}

	err = yaml.Unmarshal(content, &tmp)
	if err != nil {
		return "", err
	}
	result, err := yaml.Marshal(tmp)
	if err != nil {
		return "", err
	}

	return string(result), nil
}

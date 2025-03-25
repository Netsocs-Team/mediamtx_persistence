package main

import (
	"encoding/json"

	"gopkg.in/yaml.v2"
)

func JSONToYAML(jsonStr string) (string, error) {
	var jsonObj any
	err := json.Unmarshal([]byte(jsonStr), &jsonObj)
	if err != nil {
		return "", err
	}

	yamlBytes, err := yaml.Marshal(jsonObj)
	if err != nil {
		return "", err
	}

	return string(yamlBytes), nil
}

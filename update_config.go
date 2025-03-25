package main

import (
	"os"

	"gopkg.in/yaml.v2"
)

func UpdateConfig(mediamtxConfiPath string, newConfig *Config) error {
	yamlContent, err := yaml.Marshal(newConfig)
	if err != nil {
		return err
	}

	err = os.WriteFile(mediamtxConfiPath+".new", yamlContent, 0644)
	if err != nil {
		return err
	}

	err = os.Rename(mediamtxConfiPath+".new", mediamtxConfiPath)
	if err != nil {
		return err
	}
	return nil
}

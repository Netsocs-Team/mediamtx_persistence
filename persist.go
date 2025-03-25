package main

import (
	"encoding/json"

	"github.com/valyala/fastjson"
)

func Persist(mediamtxAPI string) (string, error) {
	paths, err := GetGlobalPaths(mediamtxAPI, "admin", "admin")
	if err != nil {
		return "", err
	}

	base, err := GetBaseConfig(mediamtxAPI, "admin", "admin")
	if err != nil {
		return "", err
	}

	baseConfig, err := fastjson.Parse(base)
	if err != nil {
		return "", err
	}
	pathsAsJson, err := json.Marshal(paths)
	if err != nil {
		return "", err
	}
	pathsConfig, err := fastjson.ParseBytes(pathsAsJson)
	if err != nil {
		return "", err
	}
	baseConfig.Set("paths", pathsConfig)

	configAsJson := baseConfig.MarshalTo(nil)
	config, err := JSONToYAML(string(configAsJson))
	if err != nil {
		return "", err
	}
	return config, nil
}

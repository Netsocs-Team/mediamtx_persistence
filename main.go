package main

import (
	"flag"
	"reflect"
	"time"

	"go.uber.org/zap"
)

const MEDIAMTX_YAML_PATH = "./mediamtx.yml"

// inspired by https://github.com/dfpc-coe/media-infra/blob/main/persist.ts
func main() {
	logger, _ := zap.NewDevelopment()
	mediamtx := flag.String("mediamtx", "http://localhost:9997", "Mediamtx Server")
	configPath := flag.String("config", MEDIAMTX_YAML_PATH, "Mediamtx Server Config")
	flag.Parse()

	if *mediamtx == "" {
		logger.Fatal("Mediamtx Server is required")
		return
	}

	ticker := time.NewTicker(10 * time.Second)

	for range ticker.C {
		config, err := Persist(*mediamtx)
		if err != nil {
			logger.Error("Error persisting config", zap.Error(err))
			continue
		}
		currentConfig, err := LoadYamlConfig(*configPath)
		if err != nil {
			logger.Error("Error loading current config", zap.Error(err))
			continue
		}

		if !reflect.DeepEqual(config, currentConfig) {
			logger.Info("Config changed, updating")
			logDiffs(logger, config, currentConfig)
			UpdateConfig(*configPath, config)
			logger.Info("Config updated")
		}
	}
	logger.Warn("Ticker stopped - exiting...")
}

func logDiffs(logger *zap.Logger, config, currentConfig *Config) {
	v1 := reflect.ValueOf(*config)
	v2 := reflect.ValueOf(*currentConfig)
	typeOfConfig := v1.Type()

	for i := 0; i < v1.NumField(); i++ {
		fieldName := typeOfConfig.Field(i).Name
		value1 := v1.Field(i).Interface()
		value2 := v2.Field(i).Interface()

		if !reflect.DeepEqual(value1, value2) {
			logger.Info("Field changed", zap.String("field", fieldName), zap.Any("old", value2), zap.Any("new", value1))
		}
	}
}

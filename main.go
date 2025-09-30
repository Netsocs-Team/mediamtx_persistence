package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"time"

	"github.com/Netsocs-Team/driver.sdk_go/pkg/client"
	"github.com/Netsocs-Team/driver.sdk_go/pkg/objects"
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

	ports := []int{9996, 9997, 9998, 9999}
	KillPorts(ports)

	cmd := exec.Command("./mediamtx")
	// Optionally, redirect output to logger or os.Stdout/os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		logger.Fatal("Failed to start mediamtx binary", zap.Error(err))
		return
	}
	logger.Info("Started mediamtx binary", zap.Int("pid", cmd.Process.Pid))

	// Monitor the process in a goroutine
	go func() {
		err := cmd.Wait()
		if err != nil {
			cmd.Process.Kill()
			logger.Fatal("mediamtx process exited with error", zap.Error(err))

		} else {
			cmd.Process.Kill()
			logger.Fatal("mediamtx process exited, terminating persistence process")

		}
	}()

	go monitorConfig(logger, configPath, mediamtx)

	nc, err := client.New()

	if err != nil {
		logger.Fatal("Error creating client", zap.Error(err))
		return
	}
	obj, err := CreateObjects(nc)

	if err != nil {
		cmd.Process.Kill()
		logger.Fatal("Error creating object", zap.Error(err))
	}

	logger.Info("Created object", zap.String("object_id", obj.GetMetadata().ObjectID))

	nc.ListenConfig()
	cmd.Process.Kill()
}

func CreateObjects(nc *client.NetsocsDriverClient) (objects.RegistrableObject, error) {
	videoEngineParams := objects.NewVideoEngineObjectParams{}
	hostname := nc.GetSiteHost()
	siteid := nc.GetSiteID()
	videoEngineParams.Metadata.Name = "VideoEngine"
	videoEngineParams.Metadata.DeviceID = "1"
	videoEngineParams.Metadata.Domain = "Video_engine"
	videoEngineParams.Metadata.ObjectID = "netsocs.video_engine.site_" + siteid
	videoEngineParams.Setup = func(obj objects.VideoEngineObject, oc objects.ObjectController) error {
		err := obj.UpdateStateAttributes(map[string]string{
			"state":         "running",
			"api_port":      "9997",
			"hls_port":      "3889",
			"hostname":      hostname,
			"playback_port": "9996",
			"rtsp_port":     "8554",
			"site_id":       siteid,
			"webrtc_port":   "8889",
		})

		return err
	}

	videoEngine := objects.NewVideoEngineObject(videoEngineParams)
	return videoEngine, nc.RegisterObject(videoEngine)
}

func KillPorts(ports []int) {
	for _, port := range ports {
		cmd := exec.Command("fuser", "-k", fmt.Sprintf("%d/tcp", port))
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Error killing port %d: %v\nOutput: %s\n", port, err, output)
		} else {
			fmt.Printf("Successfully killed process on port %d\n", port)
		}
	}
}

func monitorConfig(logger *zap.Logger, configPath *string, mediamtx *string) {
	ticker := time.NewTicker(10 * time.Second)

	for range ticker.C {
		logger.Info("Checking for config changes")
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
			err = UpdateConfig(*configPath, config)
			if err != nil {
				logger.Error("Error updating config", zap.Error(err))
				continue
			}
			logger.Info("Config updated")
		} else {
			logger.Info("No changes to config")
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

//"4R/lZYuZ5JprR+TIdl0+Zf8nM+DyNwlm8KdSVemumZ5NjYKA1ihMFmfOqJlLE5lMtokPjBvKPj7QAh/qnEuyvxPoVztfjpMxe17JinRm/T08inFOQfyGJGPu2LPM7lceAIgZ60Iqh98M22DgjERaQMdOlU0k3f8MRgVYhYnyIAYPh35pZq6SVjsl07La3qVLB4m9Vh9mFNPwYvan5IueaXWxLCVPxSJ80/KMTcSsCoUaMJhYtcsRPU3ZyCYySCrfdK/FQqK7VRy0V7S+Hk5Lgs0jEJObEnUEsiDndqI/JG8yl5Gm45oQD/0BZTWOEtZL+yXTKXEvqi6W9X9svg/eRi1+xB2bvbQQfQQQVdaN0wBJZSte4r7Lnf+fJKTFotynn/mj1nNVoBWJmmZuE1g5CJtYOf2dH4+ga+GWmH8ParGrFgGok9PfnUtpYQgNzIksZDYLb53O57XVuXMZRXZoF13OrZvVMHS+kwaVSFy0o1QHnwdJSJ3hNMtidUoVw5aWngOt+8TWOl1s/Ev0NCcWnVBO5jAKVlkrzuS8cKkqH4XxQx+zerlpvL9MaFOfiBJsAkhcMULjjrl9naH+t2ViOfr/m2ei2xE5PTBN9G3WSMB4ctVzVRf4gkXq9pbHeCQtUV9b+rhm2KZ7ZJ9O0w8aPzRoQnNHciwRo018T6gDzGvenaIbyG0cU94ryw6o+b8vgjDqV5dg0ovbDJCgtqt+NMHdXl/Rrb/xddXNL8pn6W3FLyHAVHeO5Fi8mCx7ZQM5DzIQZJK+4OgJEJmY9ukl/x8pMp+rkq7gkXqNHOQxtsICWcVZeBcINqRI9JcvYsoDsWD+iLZvoksSZGpe66RiZeYfmC4J1hXaqwk6g3Xzi1aWOSb2GhTIuOmvCXEp9fg5scQf96Bek9Lg7/bHYylhwVILB+nZaftTwK9AB9ywwNIMYm1kk4QulFfxLT/bhioTRqw2FJW0fe1DMmEaQBw="

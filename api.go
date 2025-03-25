package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

func GetBaseConfig(mediamtxAPI string, user string, pass string) (*Config, error) {
	resp, err := resty.New().R().
		SetBasicAuth(user, pass).
		Get(fmt.Sprintf("%s/v3/config/global/get", mediamtxAPI))
	if err != nil {
		return nil, err
	}
	result := Config{}
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func GetGlobalPaths(mediamtxAPI string, user string, pass string) ([]Path, error) {
	client := resty.New()
	var paths []Path
	total := 0
	page := -1

	for {
		page++
		resp, err := client.R().
			SetBasicAuth(user, pass).
			SetQueryParams(map[string]string{
				"itemsPerPage": "1000",
				"page":         fmt.Sprintf("%d", page),
			}).
			Get(fmt.Sprintf("%s/v3/config/paths/list", mediamtxAPI))
		if err != nil {
			return nil, err
		}

		var pathResponse PathResponse
		err = json.Unmarshal(resp.Body(), &pathResponse)
		if err != nil {
			return nil, err
		}

		total = pathResponse.ItemCount

		paths = append(paths, pathResponse.Items...)

		if total <= page*1000 {
			break
		}
	}

	return paths, nil
}

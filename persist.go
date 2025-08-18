package main

func Persist(mediamtxAPI string) (*Config, error) {
	paths, err := GetGlobalPaths(mediamtxAPI, "admin", "admin")
	if err != nil {
		return nil, err
	}

	config, err := GetBaseConfig(mediamtxAPI, "admin", "admin")
	if err != nil {
		return nil, err
	}

	config.Paths = make(map[string]Path)

	for _, path := range paths {
		// path.SourceOnDemand = true
		path.RecordDeleteAfter = "10m"
		path.RecordSegmentDuration = "5m"
		config.Paths[path.Name] = path
	}

	return config, nil
}

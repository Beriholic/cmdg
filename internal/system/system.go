package system

type SystemInfo struct {
	Name    string
	Version string
}

func GetSystemInfo() (*SystemInfo, error) {
	data, err := fastfetch()
	if err != nil {
		return nil, err
	}

	return &SystemInfo{
		Name:    data.Result.Name,
		Version: data.Result.Version,
	}, nil
}

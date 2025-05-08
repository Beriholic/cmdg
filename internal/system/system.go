package system

import "log"

type SystemInfo struct {
	Name    string
	Version string
}

func GetSystemInfo() *SystemInfo {
	data, err := fastfetch()
	if err != nil {
		log.Fatal(err)
	}

	return &SystemInfo{
		Name:    data.Result.Name,
		Version: data.Result.Version,
	}
}

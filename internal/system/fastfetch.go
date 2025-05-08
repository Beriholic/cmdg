package system

import (
	"bytes"
	"encoding/json"
	"os/exec"
)

type osInfo struct {
	Result struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	} `json:"result"`
}

func fastfetch() (*osInfo, error) {
	cmd := exec.Command("fastfetch", "-s", "OS", "--format", "json")
	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf

	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	jsonStr := stdoutBuf.String()

	var data []*osInfo
	json.Unmarshal([]byte(jsonStr), &data)

	return data[0], nil
}

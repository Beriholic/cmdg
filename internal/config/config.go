package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/charmbracelet/huh"
	"github.com/spf13/viper"
)

const configFilePath = "$HOME/.config/cmdg/config.toml"

type CmdgConfig struct {
	Key   string
	Model string
}

func Verify() error {
	cfg := Get()
	if cfg.Key == "" {
		return fmt.Errorf("api key must be set, use `cmdg config` to set it")
	}
	if cfg.Model == "" {
		return fmt.Errorf("model must be set, use `cmdg config` to set it")
	}
	return nil
}

var (
	geminicConfigOnce sync.Once
	geminicConfig     *CmdgConfig = nil
)

func Get() *CmdgConfig {
	geminicConfigOnce.Do(func() {
		geminicConfig = load()
	})

	return geminicConfig
}

func Create() error {
	expandedPath := os.ExpandEnv(configFilePath)

	configDir := filepath.Dir(expandedPath)
	if err := os.MkdirAll(configDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create config directory: %v", err)
	}

	viper.SetConfigFile(expandedPath)
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to red config file: %v", err)
	}

	keyState := viper.GetString("key")
	modelState := viper.GetString("model")

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("What is your Gemini API key?").
				Value(&keyState),
			huh.NewInput().
				Title("Which model do you want to use?").
				Value(&modelState),
		).WithTheme(huh.ThemeBase()),
	)

	if err := form.Run(); err != nil {
		return fmt.Errorf("failed to get user input: %v", err)
	}

	viper.Set("key", keyState)
	viper.Set("model", modelState)

	if err := viper.WriteConfigAs(expandedPath); err != nil {
		return fmt.Errorf("failed to write config file: %v", err)
	}

	fmt.Printf("Configuration saved to %s\n", expandedPath)
	return nil
}

func load() *CmdgConfig {
	expandedPath := os.ExpandEnv(configFilePath)

	viper.SetConfigFile(expandedPath)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Failed to read config file: %v\n", err)
		return nil
	}

	return &CmdgConfig{
		Key:   viper.GetString("key"),
		Model: viper.GetString("model"),
	}
}
func SetModel(model string) error {
	expandedPath := os.ExpandEnv(configFilePath)
	viper.SetConfigFile(expandedPath)

	model = strings.TrimPrefix(model, "models/")

	viper.Set("model", model)

	if err := viper.WriteConfigAs(expandedPath); err != nil {
		return fmt.Errorf("failed to write config file: %v", err)
	}
	fmt.Printf("Configuration saved to %s\n", expandedPath)
	return nil
}

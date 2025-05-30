package internal

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/beriholic/cmdg/internal/config"
	"github.com/beriholic/cmdg/internal/service"
	"github.com/beriholic/cmdg/internal/ui"
	"github.com/charmbracelet/huh"
)

func GeneratorCommand(ctx context.Context, input string) error {
	gs, err := service.NewGeminiServer(ctx, input)
	if err != nil {
		return err
	}
	res, err := gs.Generate(ctx)

	if err != nil {
		return err
	}

	fmt.Println(strings.Join(res.Cmds, "\n"))

	ExecutorCommand(res.Cmds)

	return nil
}

func ExecutorCommand(cmds []string) error {
	comfirmValue := false

	confirm := huh.NewConfirm().
		Title("Confirm to run?").
		Affirmative("Yes").
		Negative("No").
		Value(&comfirmValue).
		WithTheme(huh.ThemeBase())

	if err := confirm.Run(); err != nil {
		return err
	}

	if !comfirmValue {
		return nil
	}

	for _, c := range cmds {
		var command *exec.Cmd
		if runtime.GOOS == "windows" {
			command = exec.Command("cmd", "/C", c)
		} else {
			command = exec.Command("bash", "-c", c)
		}

		command.Stdout = os.Stdout
		command.Stderr = os.Stderr

		err := command.Run()
		if err != nil {
			return err
		}
	}

	return nil
}

func UpdateGeminiModelSelect(ctx context.Context) error {
	geminiService, err := service.NewGeminiServer(ctx, "")
	if err != nil {
		return err
	}
	models := geminiService.ListModels(ctx)
	model, err := ui.RenderStringsSelect(models)
	if err != nil {
		return err
	}

	if err = config.SetModel(model); err != nil {
		return err
	}

	return nil
}

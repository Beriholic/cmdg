package internal

import (
	"context"
	"fmt"

	"github.com/beriholic/cmdg/internal/system"
)

func GeneratorCommand(ctx context.Context, input string) error {
	system, err := system.GetSystemInfo()

	if err != nil {
		return err
	}
	fmt.Println(system)
	return nil

}

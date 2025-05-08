package internal

import (
	"context"
	"fmt"

	"github.com/beriholic/cmdg/internal/service"
)

func GeneratorCommand(ctx context.Context, input string) error {
	prompt := service.NewPrompt().Build(input)
	fmt.Println(prompt)
	return nil
}

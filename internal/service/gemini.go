package service

import (
	"context"

	"github.com/beriholic/cmdg/internal/config"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type GeminiServer struct {
	Prompt string
	client *genai.Client
}

func NewGeminiServer(ctx context.Context, input string) (*GeminiServer, error) {
	prompt := NewPrompt().Build(input)

	client, err := genai.NewClient(
		ctx,
		option.WithAPIKey(config.Get().Key),
	)
	if err != nil {
		return nil, err
	}

	return &GeminiServer{
		Prompt: prompt,
		client: client,
	}, nil
}

func (g *GeminiServer) Generate() {

}
func (g *GeminiServer) CloseClient() {
	g.client.Close()
}

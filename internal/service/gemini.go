package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/beriholic/cmdg/internal/config"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type GeminiServer struct {
	Prompt string
	client *genai.Client
}
type GenerateResult struct {
	Cmds []string `json:"cmd"`
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

func (g *GeminiServer) Generate(ctx context.Context) (*GenerateResult, error) {
	model := g.client.GenerativeModel(config.Get().Model)

	model.SetTemperature(0.9)
	model.ResponseMIMEType = "application/json"

	resp, err := model.GenerateContent(ctx, genai.Text(g.Prompt))
	if err != nil {
		return nil, err
	}

	jsonStr := fmt.Sprintf("%s", resp.Candidates[0].Content.Parts[0])

	var res *GenerateResult
	json.Unmarshal([]byte(jsonStr), &res)

	return res, nil
}

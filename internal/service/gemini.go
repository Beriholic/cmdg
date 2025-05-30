package service

import (
	"context"
	"encoding/json"

	"github.com/beriholic/cmdg/internal/config"
	"google.golang.org/genai"
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

	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  config.Get().Key,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return nil, err
	}

	return &GeminiServer{
		Prompt: prompt,
		client: client,
	}, nil
}

func (g *GeminiServer) Generate(ctx context.Context) (*GenerateResult, error) {
	geminiConfig := &genai.GenerateContentConfig{
		ResponseMIMEType: "application/json",
		ResponseSchema: &genai.Schema{
			Type: genai.TypeObject,
			Properties: map[string]*genai.Schema{
				"cmd": {
					Type: genai.TypeArray,
					Items: &genai.Schema{
						Type: genai.TypeString,
					},
				},
			},
		},
	}

	result, err := g.client.Models.GenerateContent(ctx, config.Get().Model, genai.Text(g.Prompt), geminiConfig)
	if err != nil {
		return nil, err
	}

	var res *GenerateResult
	json.Unmarshal([]byte(result.Text()), &res)

	return res, nil
}

func (g *GeminiServer) ListModels(ctx context.Context) []string {
	iter := g.client.Models.All(ctx)

	models := []string{}

	for model, err := range iter {
		if err != nil {
			continue
		}

		models = append(models, model.Name)
	}

	return models
}

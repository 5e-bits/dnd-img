package prompt

import (
	"context"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
)

// Generator defines the interface for prompt generation
type Generator interface {
	Generate(ctx context.Context, subject string) (string, error)
}

// OpenAIGenerator implements the Generator interface using OpenAI's API
type OpenAIGenerator struct {
	client          *openai.Client
	systemPrompt    string
	webSearchPrompt string
}

// NewOpenAIGenerator creates a new OpenAI prompt generator
func NewOpenAIGenerator(client *openai.Client, systemPrompt string, webSearchPrompt string) *OpenAIGenerator {
	return &OpenAIGenerator{
		client:          client,
		systemPrompt:    systemPrompt,
		webSearchPrompt: webSearchPrompt,
	}
}

// Generate creates a prompt using OpenAI's API
func (g *OpenAIGenerator) Generate(ctx context.Context, subject string) (string, error) {
	resp, err := g.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT4o,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: g.systemPrompt + "\nGenerate a concise image prompt under 1000 characters that captures the essence of the subject.",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: g.webSearchPrompt + "\nSubject: " + subject,
				},
			},
		},
	)

	if err != nil {
		return "", fmt.Errorf("failed to generate prompt: %w", err)
	}

	// Ensure the prompt is under 1000 characters
	prompt := resp.Choices[0].Message.Content
	if len(prompt) > 1000 {
		prompt = prompt[:1000]
	}

	return prompt, nil
}

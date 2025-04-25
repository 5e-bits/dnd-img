package prompt

import (
	"context"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
)

const (
	PROMPT_SYSTEM     = `You are an expert in Dungeons & Dragons art direction, specializing in the iconic 90s TSR catalog style. Your task is to create focused, clear descriptions that depict individual D&D subjects (items, monsters, weapons, spells, etc.) in a way that would fit perfectly in a D&D catalog or rulebook. Focus on clear, isolated depictions that showcase the subject's key features while maintaining the dramatic, high-fantasy aesthetic of classic TSR-era artwork.`
	PROMPT_WEB_SEARCH = `Research and analyze this D&D subject, focusing on:
1. Core visual characteristics and defining features
2. Historical context and significance in D&D lore
3. Typical appearance and key details
4. Common interactions or effects (if applicable)
5. Key artistic elements from 90s D&D catalog style (clear composition, dramatic lighting, rich textures)

Generate a detailed description that would serve as a perfect prompt for creating a piece of art that could have appeared in a 90s D&D catalog or rulebook.`
)

// Generator defines the interface for prompt generation
type Generator interface {
	Generate(ctx context.Context, subject string) (string, error)
}

// OpenAIGenerator implements the Generator interface using OpenAI's API
type OpenAIGenerator struct {
	client       *openai.Client
	systemPrompt string
}

// NewOpenAIGenerator creates a new OpenAI prompt generator
func NewOpenAIGenerator(client *openai.Client, systemPrompt string) *OpenAIGenerator {
	return &OpenAIGenerator{
		client:       client,
		systemPrompt: systemPrompt,
	}
}

// Generate creates a prompt using OpenAI's API
func (g *OpenAIGenerator) Generate(ctx context.Context, subject string) (string, error) {
	resp, err := g.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT4oMini20240718,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: g.systemPrompt + "\nGenerate a concise image prompt under 1000 characters that captures the essence of the subject.",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: PROMPT_WEB_SEARCH + "\nSubject: " + subject,
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

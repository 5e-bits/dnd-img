package image

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"image/png"
	"os"
	"path/filepath"

	"github.com/charmbracelet/log"
	openai "github.com/sashabaranov/go-openai"
)

const (
	PROMPT_IMAGE_GENERATION = `Create a high-quality digital painting in the style of 90s D&D catalog art with these specifications:
- Aspect ratio: 1:1
- Style: Hand-painted, high-fantasy illustration
- Subject: Clear, isolated depiction of the requested subject
- Lighting: Dramatic, atmospheric lighting with strong contrast
- Composition: Centered, clear presentation of the subject
- Details: Rich textures, intricate details in the subject
- Color palette: Rich, saturated colors typical of 90s fantasy art
- Background: Simple, dark or atmospheric background that doesn't distract from the subject
- Mood: Mysterious, powerful, or iconic depending on the subject
- Quality: Professional illustration quality suitable for a D&D catalog`
)

// Generator defines the interface for image generation
type Generator interface {
	Generate(ctx context.Context, prompt string) (string, error)
	SaveImage(base64Data, filename string) error
}

// OpenAIGenerator implements the Generator interface using OpenAI's API
type OpenAIGenerator struct {
	client *openai.Client
}

// NewOpenAIGenerator creates a new OpenAI image generator
func NewOpenAIGenerator(client *openai.Client) *OpenAIGenerator {
	return &OpenAIGenerator{
		client: client,
	}
}

// Generate creates an image using OpenAI's API
func (g *OpenAIGenerator) Generate(ctx context.Context, prompt string) (string, error) {
	req := openai.ImageRequest{
		Prompt:         prompt,
		Model:          openai.CreateImageModelDallE3,
		Size:           openai.CreateImageSize1024x1024,
		ResponseFormat: openai.CreateImageResponseFormatB64JSON,
		N:              1,
	}

	resp, err := g.client.CreateImage(ctx, req)
	if err != nil {
		return "", fmt.Errorf("failed to create image: %w", err)
	}

	return resp.Data[0].B64JSON, nil
}

// SaveImage saves the base64 encoded image to a file
func (g *OpenAIGenerator) SaveImage(base64Data, filename string) error {
	imgBytes, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return fmt.Errorf("failed to decode image: %w", err)
	}

	img, err := png.Decode(bytes.NewReader(imgBytes))
	if err != nil {
		return fmt.Errorf("failed to decode PNG: %w", err)
	}

	// Ensure output directory exists
	if err := os.MkdirAll("output", 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	filepath := filepath.Join("output", filename)
	file, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	if err := png.Encode(file, img); err != nil {
		return fmt.Errorf("failed to save image: %w", err)
	}

	log.Info("Image saved", "filename", filepath)
	return nil
}

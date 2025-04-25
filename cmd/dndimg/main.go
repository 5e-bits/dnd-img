package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/alecthomas/kong"
	"github.com/charmbracelet/log"
	"github.com/sashabaranov/go-openai"

	"github.com/5e-bits/dndimg/internal/config"
	"github.com/5e-bits/dndimg/internal/image"
	"github.com/5e-bits/dndimg/internal/prompt"
)

type CLI struct {
	Subject string `arg:"" help:"The D&D subject to generate an image for"`
}

func run() error {
	var cli CLI
	kong.Parse(&cli,
		kong.Name("dndimg"),
		kong.Description("Generate D&D SRD subject images using AI"),
		kong.UsageOnError(),
	)

	// Initialize configuration
	cfg := config.New()

	// Initialize OpenAI client
	client := openai.NewClient(cfg.OpenAIToken)
	ctx := context.Background()

	// Initialize generators
	promptGen := prompt.NewOpenAIGenerator(client, cfg.SystemPrompt)
	imageGen := image.NewOpenAIGenerator(client)

	// Step 1: Generate prompt
	log.Info("Generating detailed description...")
	prompt, err := promptGen.Generate(ctx, cli.Subject)
	if err != nil {
		return fmt.Errorf("failed to generate prompt: %w", err)
	}

	// Step 2: Generate image
	log.Info("Creating image...")
	imageData, err := imageGen.Generate(ctx, prompt)
	if err != nil {
		return fmt.Errorf("failed to generate image: %w", err)
	}

	// Step 3: Save image
	log.Info("Saving image...")
	filename := fmt.Sprintf("%s.png", formatFilename(cli.Subject))
	if err := imageGen.SaveImage(imageData, filename); err != nil {
		return fmt.Errorf("failed to save image: %w", err)
	}

	return nil
}

// formatFilename formats a subject name into a valid filename
func formatFilename(name string) string {
	// Convert to lowercase
	name = strings.ToLower(name)
	// Replace spaces with dashes
	name = strings.ReplaceAll(name, " ", "-")
	return name
}

func main() {
	log.SetLevel(log.DebugLevel)
	log.SetTimeFormat(time.Kitchen)
	log.SetPrefix("dndimg")

	if err := run(); err != nil {
		log.Fatal("Error", "error", err)
	}
}

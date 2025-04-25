package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
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
	Subject      string `arg:"" optional:"" help:"The D&D subject to generate an image for"`
	SubjectsFile string `short:"f" help:"Path to a file containing newline-delimited subjects"`
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
	promptGen := prompt.NewOpenAIGenerator(client, cfg.SystemPrompt, cfg.WebSearchPrompt)
	imageGen := image.NewOpenAIGenerator(client)

	// Create a rate limiter (1 request per 20 seconds to be safe)
	rateLimiter := time.Tick(20 * time.Second)

	// Function to process a single subject
	processSubject := func(subject string) error {
		// Step 1: Generate prompt
		log.Info("Generating detailed description...", "subject", subject)
		prompt, err := promptGen.Generate(ctx, subject)
		if err != nil {
			return fmt.Errorf("failed to generate prompt for %s: %w", subject, err)
		}

		// Step 2: Generate image
		log.Info("Creating image...", "subject", subject)
		imageData, err := imageGen.Generate(ctx, prompt)
		if err != nil {
			return fmt.Errorf("failed to generate image for %s: %w", subject, err)
		}

		// Step 3: Save image
		log.Info("Saving image...", "subject", subject)
		filename := fmt.Sprintf("%s.png", formatFilename(subject))
		if err := imageGen.SaveImage(imageData, filename); err != nil {
			return fmt.Errorf("failed to save image for %s: %w", subject, err)
		}

		return nil
	}

	// Handle single subject case
	if cli.Subject != "" {
		return processSubject(cli.Subject)
	}

	// Handle subjects file case
	if cli.SubjectsFile == "" {
		return fmt.Errorf("either a subject or subjects file must be provided")
	}

	// Read subjects from file
	file, err := os.Open(cli.SubjectsFile)
	if err != nil {
		return fmt.Errorf("failed to open subjects file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalSubjects := 0
	successfulSubjects := 0

	// Count total subjects first
	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) != "" {
			totalSubjects++
		}
	}

	// Reset file pointer
	file.Seek(0, 0)
	scanner = bufio.NewScanner(file)

	log.Info("Starting batch processing", "total_subjects", totalSubjects)

	// Process each subject
	for scanner.Scan() {
		subject := strings.TrimSpace(scanner.Text())
		if subject == "" {
			continue
		}

		// Wait for rate limiter
		<-rateLimiter

		if err := processSubject(subject); err != nil {
			log.Error("Failed to process subject", "subject", subject, "error", err)
			continue
		}

		successfulSubjects++
		log.Info("Progress", "completed", successfulSubjects, "total", totalSubjects)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading subjects file: %w", err)
	}

	log.Info("Batch processing complete",
		"successful", successfulSubjects,
		"total", totalSubjects,
		"failed", totalSubjects-successfulSubjects)

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

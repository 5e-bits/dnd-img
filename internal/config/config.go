package config

import (
	"os"

	"github.com/5e-bits/dndimg/internal/prompt"
	"github.com/charmbracelet/log"
	_ "github.com/joho/godotenv/autoload"
)

// Config holds the application configuration
type Config struct {
	OpenAIToken  string
	SystemPrompt string
}

// New creates a new configuration from environment variables
func New() *Config {
	token := os.Getenv("OPEN_AI_TOKEN")
	if token == "" {
		log.Fatal("OPEN_AI_TOKEN environment variable is required")
	}

	systemPrompt := os.Getenv("SYSTEM_PROMPT")
	if systemPrompt == "" {
		systemPrompt = prompt.PROMPT_SYSTEM
	}

	return &Config{
		OpenAIToken:  token,
		SystemPrompt: systemPrompt,
	}
}

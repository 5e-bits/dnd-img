package config

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/charmbracelet/log"
	_ "github.com/joho/godotenv/autoload"
)

// Config holds the application configuration
type Config struct {
	OpenAIToken       string
	SystemPrompt      string
	WebSearchPrompt   string
	SubjectsDelimiter string
}

const (
	PROMPT_SYSTEM     = `You are an expert in Dungeons & Dragons art direction, specializing in the iconic 90s TSR catalog style. Your task is to create focused, clear descriptions that depict individual D&D subjects (items, monsters, weapons, spells, etc.) in a way that would fit perfectly in a D&D catalog or rulebook. Focus on clear, isolated depictions that showcase the subject's key features while maintaining the dramatic, high-fantasy aesthetic of classic TSR-era hand-painted artwork. Ensure the descriptions emphasize traditional painting techniques, brush strokes, and artistic style rather than photographic realism.`
	PROMPT_WEB_SEARCH = `Research and analyze this D&D subject, focusing on:
1. Core visual characteristics and defining features
2. Historical context and significance in D&D lore
3. Typical appearance and key details
4. Common interactions or effects (if applicable)
5. Key artistic elements from 90s D&D catalog style (clear composition, dramatic lighting, rich textures, hand-painted aesthetic)

Generate a detailed description that would serve as a perfect prompt for creating a piece of hand-painted art that could have appeared in a 90s D&D catalog or rulebook.`
	DEFAULT_DELIMITER = "\n"
)

// New creates a new configuration from environment variables
func New() *Config {
	token := os.Getenv("OPEN_AI_TOKEN")
	if token == "" {
		log.Fatal("OPEN_AI_TOKEN environment variable is required")
	}

	systemPrompt := os.Getenv("SYSTEM_PROMPT")
	if systemPrompt == "" {
		systemPrompt = PROMPT_SYSTEM
	}

	webSearchPrompt := os.Getenv("WEB_SEARCH_PROMPT")
	if webSearchPrompt == "" {
		webSearchPrompt = PROMPT_WEB_SEARCH
	}

	delimiter := os.Getenv("SUBJECTS_DELIMITER")
	if delimiter == "" {
		delimiter = DEFAULT_DELIMITER
	}

	return &Config{
		OpenAIToken:       token,
		SystemPrompt:      systemPrompt,
		WebSearchPrompt:   webSearchPrompt,
		SubjectsDelimiter: delimiter,
	}
}

// ProcessSubjectsFile reads subjects from a file and returns them as a slice
func (c *Config) ProcessSubjectsFile(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to open subjects file: %w", err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read subjects file: %w", err)
	}

	// Split content by the configured delimiter
	subjects := strings.Split(string(content), c.SubjectsDelimiter)

	// Filter out empty lines and trim whitespace
	var filteredSubjects []string
	for _, subject := range subjects {
		subject = strings.TrimSpace(subject)
		if subject != "" {
			filteredSubjects = append(filteredSubjects, subject)
		}
	}

	return filteredSubjects, nil
}

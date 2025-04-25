# dndimg

A command-line tool for generating D&D SRD subject images using AI. This tool creates high-quality, stylized images of D&D subjects (items, monsters, weapons, spells, etc.) in the iconic 90s TSR catalog style.

## Features

- Generate detailed image prompts using AI
- Create high-quality D&D subject images
- Save images in PNG format
- Command-line interface for easy use
- Configurable through environment variables
- Batch processing of multiple subjects
- Rate limiting to respect API limits

## Installation

### Prerequisites

- Go 1.23 or later
- OpenAI API key

### Using Go

```bash
go install github.com/5e-bits/dndimg@latest
```

### From Source

```bash
git clone https://github.com/5e-bits/dndimg.git
cd dndimg
go install
```

## Configuration

Create a `.env` file in your working directory with the following variables:

```env
OPENAI_API_KEY=your_api_key_here
```

### System Prompt

The tool uses a default system prompt that guides the AI in generating D&D-themed images in the 90s TSR catalog style. You can override this by setting the `SYSTEM_PROMPT` environment variable.

Default system prompt:
```text
You are an expert in Dungeons & Dragons art direction, specializing in the iconic 90s TSR catalog style. Your task is to create focused, clear descriptions that depict individual D&D subjects (items, monsters, weapons, spells, etc.) in a way that would fit perfectly in a D&D catalog or rulebook. Focus on clear, isolated depictions that showcase the subject's key features while maintaining the dramatic, high-fantasy aesthetic of classic TSR-era hand-painted artwork. Ensure the descriptions emphasize traditional painting techniques, brush strokes, and artistic style rather than photographic realism.
```

### Web Search Prompt

The tool also uses a default web search prompt to guide the AI in researching and analyzing D&D subjects. You can override this by setting the `WEB_SEARCH_PROMPT` environment variable.

Default web search prompt:
```text
Research and analyze this D&D subject, focusing on:
1. Core visual characteristics and defining features
2. Historical context and significance in D&D lore
3. Typical appearance and key details
4. Common interactions or effects (if applicable)
5. Key artistic elements from 90s D&D catalog style (clear composition, dramatic lighting, rich textures, hand-painted aesthetic)

Generate a detailed description that would serve as a perfect prompt for creating a piece of hand-painted art that could have appeared in a 90s D&D catalog or rulebook.
```

To use custom prompts, add them to your `.env` file:
```env
SYSTEM_PROMPT=your_custom_system_prompt_here
WEB_SEARCH_PROMPT=your_custom_web_search_prompt_here
```

## Usage

### Single Subject Mode

Generate an image for any D&D subject:

```bash
dndimg "Ancient Red Dragon"
```

The tool will:
1. Generate a detailed description of the subject
2. Create an image based on the description
3. Save the image as a PNG file in the current directory

### Batch Processing Mode

Process multiple subjects from a file:

```bash
dndimg -f subjects.txt
```

Where `subjects.txt` contains one subject per line:
```
Ancient Red Dragon
Beholder
Mind Flayer
```

The tool will:
1. Process each subject in sequence
2. Respect API rate limits (1 request per 20 seconds)
3. Show progress for each subject
4. Continue processing if any subject fails
5. Provide a summary of successful and failed subjects

## Examples

```bash
# Generate an image of a weapon
dndimg "Heavy crossbow"

# Generate an image of a monster
dndimg "Brown Bear"

# Generate an image of a spell
dndimg "Chill touch"

# Process multiple subjects from a file
dndimg -f subjects.txt
```

## Development

### Dependencies

- [github.com/alecthomas/kong](https://github.com/alecthomas/kong) - Command-line argument parsing
- [github.com/charmbracelet/log](https://github.com/charmbracelet/log) - Logging
- [github.com/sashabaranov/go-openai](https://github.com/sashabaranov/go-openai) - OpenAI API client

### Building

```bash
go build -o dndimg cmd/dndimg/main.go
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request. 
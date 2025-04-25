# dndimg

A command-line tool for generating D&D SRD subject images using AI. This tool creates high-quality, stylized images of D&D subjects (items, monsters, weapons, spells, etc.) in the iconic 90s TSR catalog style.

## Features

- Generate detailed image prompts using AI
- Create high-quality D&D subject images
- Save images in PNG format
- Command-line interface for easy use
- Configurable through environment variables

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
You are an expert in Dungeons & Dragons art direction, specializing in the iconic 90s TSR catalog style. Your task is to create focused, clear descriptions that depict individual D&D subjects (items, monsters, weapons, spells, etc.) in a way that would fit perfectly in a D&D catalog or rulebook. Focus on clear, isolated depictions that showcase the subject's key features while maintaining the dramatic, high-fantasy aesthetic of classic TSR-era artwork.
```

To use a custom prompt, add it to your `.env` file:
```env
SYSTEM_PROMPT=your_custom_prompt_here
```

## Usage

Generate an image for any D&D subject:

```bash
dndimg "Ancient Red Dragon"
```

The tool will:
1. Generate a detailed description of the subject
2. Create an image based on the description
3. Save the image as a PNG file in the current directory

## Examples

```bash
# Generate an image of a weapon
dndimg "Heavy crossbow"

# Generate an image of a monster
dndimg "Brown Bear"

# Generate an image of a spell
dndimg "Chill touch"
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
# sai (Stream AI)

A command-line utility that processes input streams through OpenAI's API, similar to `sed` but for AI processing. Written in Go for maximum performance and easy distribution.

## Table of Contents

- [Overview](#overview)
- [Installation](#installation)
- [Quick Start](#quick-start)
- [Command-Line API Reference](#command-line-api-reference)
- [Usage Examples](#usage-examples)
- [Configuration](#configuration)
- [Error Handling](#error-handling)
- [Development](#development)
- [API Architecture](#api-architecture)
- [Performance](#performance)
- [Contributing](#contributing)
- [License](#license)

## Overview

`sai` is a streaming AI processor that acts as a bridge between your command-line workflows and OpenAI's powerful language models. It reads input from stdin, processes it through OpenAI's API, and streams the response back to stdout in real-time.

### Key Features

- **Streaming responses** - Real-time output as the AI generates content
- **Zero dependencies** - Single static binary with no runtime requirements
- **Cross-platform** - Native binaries for macOS, Linux, and Windows
- **Pipeline friendly** - Designed for Unix-style command chaining
- **Flexible context** - Add custom context to guide AI responses
- **Code-focused mode** - Clean code output without explanations
- **Multiple models** - Support for different OpenAI models

## Installation

### Pre-built Binaries (Recommended)

Download the latest release for your platform:

```bash
# macOS (Intel)
curl -L https://github.com/rwejlgaard/sai/releases/latest/download/sai-darwin-amd64 -o sai
chmod +x sai && sudo mv sai /usr/local/bin/

# macOS (Apple Silicon)
curl -L https://github.com/rwejlgaard/sai/releases/latest/download/sai-darwin-arm64 -o sai
chmod +x sai && sudo mv sai /usr/local/bin/

# Linux (x86_64)
curl -L https://github.com/rwejlgaard/sai/releases/latest/download/sai-linux-amd64 -o sai
chmod +x sai && sudo mv sai /usr/local/bin/

# Linux (ARM64)
curl -L https://github.com/rwejlgaard/sai/releases/latest/download/sai-linux-arm64 -o sai
chmod +x sai && sudo mv sai /usr/local/bin/
```

For Windows, download the appropriate `.exe` file from the [releases page](https://github.com/rwejlgaard/sai/releases/latest) and add it to your PATH.

### Build from Source

```bash
# Clone the repository
git clone https://github.com/rwejlgaard/sai.git
cd sai

# Build and install
make build
sudo make install

# Or install directly with Go
go install github.com/rwejlgaard/sai/src/cmd/sai@latest
```

## Quick Start

1. **Set up your API key**:
   ```bash
   export OPENAI_API_KEY='your-api-key-here'
   ```

2. **Basic usage**:
   ```bash
   echo "Hello, world!" | sai
   ```

3. **With context**:
   ```bash
   echo "Paris" | sai --context "List 3 famous landmarks in"
   ```

4. **Generate code**:
   ```bash
   echo "Python function to calculate fibonacci" | sai --code
   ```

## Command-Line API Reference

### Synopsis

```bash
sai [OPTIONS]
```

### Options

| Flag | Short | Type | Default | Description |
|------|-------|------|---------|-------------|
| `--context` | `-c` | string | `""` | Context to prepend to the input |
| `--model` | `-m` | string | `"gpt-4o-mini"` | OpenAI model to use |
| `--api-key` | `-k` | string | `""` | OpenAI API key (overrides environment variable) |
| `--code` | `-o` | bool | `false` | Return only code without explanations or backticks |

### Input/Output

- **Input**: Reads from stdin (pipe or redirect)
- **Output**: Streams response to stdout in real-time
- **Errors**: Written to stderr with descriptive messages

### Exit Codes

- `0`: Success
- `1`: Error (missing API key, no input, API error, network issue, etc.)

## Usage Examples

### Basic Usage

```bash
# Simple AI processing
echo "Explain quantum computing in simple terms" | sai

# Ask questions
echo "What is the capital of France?" | sai

# Process with context
echo "machine learning" | sai --context "Explain the concept of"
```

### File Processing

```bash
# Summarize a document
cat document.txt | sai --context "Summarize this document:"

# Review code
cat main.go | sai --context "Review this Go code and suggest improvements:"

# Transform data
cat data.json | sai --context "Convert this JSON to YAML format:"

# Generate documentation
cat *.go | sai --context "Generate API documentation for this Go code:"
```

### Code Generation

```bash
# Generate clean code (no explanations)
echo "REST API server in Go" | sai --code

# Generate with context
echo "user authentication" | sai --code --context "Create a secure Go function for"

# Generate tests
cat function.go | sai --code --context "Write comprehensive unit tests for:"
```

### Advanced Workflows

```bash
# Different models for different tasks
echo "Complex mathematical proof" | sai --model gpt-4
echo "Simple translation" | sai --model gpt-3.5-turbo

# Chained processing
echo "Create a REST API" | sai --code | sai --context "Add comprehensive error handling to this code:"

# Multiple transformations
cat requirements.txt | sai --context "Explain each dependency:" | sai --context "Summarize the main categories:"

# Custom API key for specific requests
echo "Sensitive query" | sai --api-key sk-your-alternative-key
```

### Integration with Other Tools

```bash
# With curl for API testing
curl -s https://api.example.com/data | sai --context "Analyze this API response:"

# With find for code analysis
find . -name "*.go" -exec cat {} \; | sai --context "Find potential security issues in this Go code:"

# With git for commit messages
git diff --cached | sai --context "Write a concise commit message for these changes:"

# With jq for JSON processing
echo '{"users": [{"name": "John", "age": 30}]}' | jq . | sai --context "Convert to XML format:"
```

## Configuration

### Environment Variables

#### `OPENAI_API_KEY` (Required)

Your OpenAI API key for authentication. Get one from [OpenAI's platform](https://platform.openai.com/api-keys).

```bash
export OPENAI_API_KEY='sk-your-api-key-here'
```

### Available Models

The tool supports all OpenAI chat models:

- `gpt-4o-mini` (default) - Cost-effective for most tasks
- `gpt-4o` - Most capable model
- `gpt-4-turbo` - High performance with large context
- `gpt-4` - Strong reasoning capabilities
- `gpt-3.5-turbo` - Fast and economical

### Configuration Best Practices

```bash
# Set API key in your shell profile
echo 'export OPENAI_API_KEY="sk-your-key"' >> ~/.bashrc

# Create aliases for common patterns
alias sai-code='sai --code'
alias sai-review='sai --context "Review this code:"'
alias sai-explain='sai --context "Explain this:"'

# Use different models for different tasks
alias sai-complex='sai --model gpt-4o'
alias sai-simple='sai --model gpt-3.5-turbo'
```

## Error Handling

### Common Error Scenarios

1. **Missing API Key**
   ```
   Error: OpenAI API key must be provided either via --api-key or OPENAI_API_KEY environment variable
   ```

2. **No Input Provided**
   ```
   Error: No input provided. Please pipe some input to sai.
   Example: echo 'Hello, world!' | sai
   ```

3. **Invalid API Key**
   ```
   Error from API (status 401): Invalid API key
   ```

4. **Network Issues**
   ```
   Error sending request: context deadline exceeded
   ```

5. **API Rate Limits**
   ```
   Error from API (status 429): Rate limit exceeded
   ```

### Error Handling Best Practices

```bash
# Check if sai is available
if ! command -v sai &> /dev/null; then
    echo "sai is not installed"
    exit 1
fi

# Verify API key is set
if [ -z "$OPENAI_API_KEY" ]; then
    echo "Please set OPENAI_API_KEY environment variable"
    exit 1
fi

# Handle errors gracefully
echo "test input" | sai || echo "Failed to process input"

# Retry with exponential backoff
for i in {1..3}; do
    echo "input" | sai && break
    sleep $((2**i))
done
```

## Development

### Build Requirements

- Go 1.22.5 or later
- No external dependencies (uses Go standard library only)

### Build Commands

```bash
# Build for current platform
make build

# Install system-wide
sudo make install

# Clean build artifacts
make clean

# Run tests
make test

# Format code
make format

# Lint code
make lint
```

### Project Structure

```
sai/
├── src/
│   └── cmd/
│       └── sai/
│           └── main.go          # Main application code
├── .github/
│   └── workflows/
│       └── release.yml          # CI/CD automation
├── Makefile                     # Build automation
├── go.mod                       # Go module definition
├── README.md                    # This file
└── .gitignore                   # Git ignore rules
```

### Cross-Platform Building

```bash
# Build for all platforms
GOOS=linux GOARCH=amd64 go build -o dist/sai-linux-amd64 ./src/cmd/sai
GOOS=linux GOARCH=arm64 go build -o dist/sai-linux-arm64 ./src/cmd/sai
GOOS=darwin GOARCH=amd64 go build -o dist/sai-darwin-amd64 ./src/cmd/sai
GOOS=darwin GOARCH=arm64 go build -o dist/sai-darwin-arm64 ./src/cmd/sai
GOOS=windows GOARCH=amd64 go build -o dist/sai-windows-amd64.exe ./src/cmd/sai
```

## API Architecture

### Core Components

1. **Command-line Parser** - Processes flags and arguments
2. **Input Processor** - Reads and validates stdin input
3. **HTTP Client** - Handles OpenAI API communication
4. **Stream Handler** - Processes real-time responses

### Data Flow

```
Input (stdin) → Context Processing → OpenAI API → Streaming Response → Output (stdout)
```

### Go Type Definitions

#### Message

```go
type Message struct {
    Role    string `json:"role"`    // "user", "assistant", or "system"
    Content string `json:"content"` // The message content
}
```

#### Request

```go
type Request struct {
    Model    string    `json:"model"`    // Model identifier
    Messages []Message `json:"messages"` // Conversation messages
    Stream   bool      `json:"stream"`   // Enable streaming
}
```

#### Response

```go
type Response struct {
    Choices []Choice `json:"choices"` // Response choices
}

type Choice struct {
    Delta Delta `json:"delta"` // Incremental content
}

type Delta struct {
    Content string `json:"content"` // Text content
}
```

### Constants

```go
const (
    defaultModel = "gpt-4o-mini"
    apiURL       = "https://api.openai.com/v1/chat/completions"
)
```

### Key Functions

#### main()

Entry point that orchestrates the entire process:

1. Parse command-line flags
2. Validate API key (flag or environment)
3. Check for stdin input availability
4. Read input from stdin with buffering
5. Construct prompt with optional context
6. Create HTTP request to OpenAI API
7. Process streaming response
8. Output results to stdout

#### Input Processing

```go
// Check if stdin has data
stat, _ := os.Stdin.Stat()
if (stat.Mode() & os.ModeCharDevice) != 0 {
    // Error: no piped input
}

// Read all input
reader := bufio.NewReader(os.Stdin)
var input strings.Builder
for {
    line, err := reader.ReadString('\n')
    if err == io.EOF {
        break
    }
    input.WriteString(line)
}
```

#### API Integration

```go
// Create request
req := Request{
    Model: *model,
    Messages: []Message{
        {Role: "user", Content: prompt},
    },
    Stream: true,
}

// Send to OpenAI
httpReq, _ := http.NewRequest("POST", apiURL, bytes.NewBuffer(reqBody))
httpReq.Header.Set("Content-Type", "application/json")
httpReq.Header.Set("Authorization", "Bearer "+*apiKey)
```

#### Streaming Response

```go
// Process server-sent events
scanner := bufio.NewScanner(resp.Body)
for scanner.Scan() {
    line := scanner.Text()
    if strings.HasPrefix(line, "data: ") {
        data := strings.TrimPrefix(line, "data: ")
        var response Response
        json.Unmarshal([]byte(data), &response)
        if len(response.Choices) > 0 {
            fmt.Print(response.Choices[0].Delta.Content)
        }
    }
}
```

## Performance

### Memory Usage

- **Streaming Processing**: Minimal memory footprint with real-time output
- **Efficient String Handling**: Uses `strings.Builder` for concatenation
- **No Caching**: Stateless operation, no persistent memory usage

### Network Efficiency

- **HTTP/1.1**: Single connection per request with keep-alive
- **Streaming**: Server-sent events for real-time responses
- **Compression**: Automatic gzip compression support

### Binary Characteristics

- **Size**: ~2-5MB static binary
- **Dependencies**: Zero runtime dependencies
- **Startup Time**: Near-instantaneous (<10ms)
- **Cross-Platform**: Native compilation for all major platforms

### Optimization Features

- **Buffered I/O**: Efficient stdin reading with `bufio.Reader`
- **Streaming Output**: Real-time response display
- **Error Recovery**: Graceful handling of partial responses
- **Resource Cleanup**: Proper connection and resource management

## Contributing

### Development Setup

1. **Clone the repository**:
   ```bash
   git clone https://github.com/rwejlgaard/sai.git
   cd sai
   ```

2. **Set up development environment**:
   ```bash
   # Install dependencies (Go 1.22.5+)
   go version
   
   # Build the project
   make build
   
   # Run tests
   make test
   ```

3. **Make changes and test**:
   ```bash
   # Format code
   make format
   
   # Lint code
   make lint
   
   # Test your changes
   echo "test" | ./dist/sai
   ```

### Adding New Features

1. **Command-line flags**: Add new flags in the flag parsing section
2. **Request modification**: Extend the `Request` type for new API parameters
3. **Response processing**: Update streaming response handling
4. **Error handling**: Add specific error cases with user-friendly messages

### Release Process

1. **Create a tag**:
   ```bash
   git tag -a v1.0.0 -m "Release v1.0.0"
   git push origin v1.0.0
   ```

2. **Automated Release**: GitHub Actions automatically builds and releases binaries for all platforms

### Code Guidelines

- Use Go standard library only
- Maintain streaming functionality
- Provide clear error messages
- Follow Go formatting conventions
- Add comments for complex logic

## License

MIT License - see LICENSE file for details.

## Links

- **GitHub Repository**: https://github.com/rwejlgaard/sai
- **Releases**: https://github.com/rwejlgaard/sai/releases
- **OpenAI API Documentation**: https://platform.openai.com/docs/api-reference
- **Issues**: https://github.com/rwejlgaard/sai/issues

---

**Made with ❤️ by [rwejlgaard](https://github.com/rwejlgaard)** 
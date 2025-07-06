# sai (Stream AI) - API Documentation

## Overview

`sai` is a command-line utility that processes input streams through OpenAI's API, similar to `sed` but for AI processing. It's designed for maximum performance and easy distribution, written in Go.

## Table of Contents

1. [Installation](#installation)
2. [Command-Line API](#command-line-api)
3. [Environment Variables](#environment-variables)
4. [Code Architecture](#code-architecture)
5. [Type Definitions](#type-definitions)
6. [Constants](#constants)
7. [Usage Examples](#usage-examples)
8. [Error Handling](#error-handling)
9. [Development Guide](#development-guide)

## Installation

### Pre-built Binaries

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

### Build from Source

```bash
# Clone and build
git clone https://github.com/rwejlgaard/sai.git
cd sai
make build

# Install system-wide
sudo make install

# Or install via Go
go install github.com/rwejlgaard/sai/src/cmd/sai@latest
```

## Command-Line API

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
- **Errors**: Written to stderr

### Exit Codes

- `0`: Success
- `1`: Error (missing API key, no input, API error, etc.)

## Environment Variables

### `OPENAI_API_KEY`

**Required** (unless `--api-key` is provided)

The OpenAI API key for authentication.

```bash
export OPENAI_API_KEY='sk-your-api-key-here'
```

## Code Architecture

### Main Components

1. **Command-line parsing** - Flag-based argument handling
2. **Input processing** - Stdin reading and context preparation
3. **OpenAI API integration** - HTTP client with streaming support
4. **Response handling** - Real-time streaming output

### Flow Diagram

```
Input (stdin) → Context Processing → OpenAI API → Streaming Response → Output (stdout)
```

## Type Definitions

### Message

Represents a single message in the OpenAI chat format.

```go
type Message struct {
    Role    string `json:"role"`    // "user", "assistant", or "system"
    Content string `json:"content"` // The message content
}
```

**Usage:**
```go
msg := Message{
    Role:    "user",
    Content: "Hello, world!",
}
```

### Request

Represents the complete request payload sent to OpenAI's API.

```go
type Request struct {
    Model    string    `json:"model"`    // Model identifier (e.g., "gpt-4o-mini")
    Messages []Message `json:"messages"` // Array of conversation messages
    Stream   bool      `json:"stream"`   // Enable streaming responses
}
```

**Usage:**
```go
req := Request{
    Model: "gpt-4o-mini",
    Messages: []Message{
        {Role: "user", Content: "Explain quantum computing"},
    },
    Stream: true,
}
```

### Delta

Represents incremental content in streaming responses.

```go
type Delta struct {
    Content string `json:"content"` // Incremental text content
}
```

### Choice

Represents a single choice/option in the API response.

```go
type Choice struct {
    Delta Delta `json:"delta"` // Incremental content for streaming
}
```

### Response

Represents the complete response structure from OpenAI's API.

```go
type Response struct {
    Choices []Choice `json:"choices"` // Array of response choices
}
```

## Constants

### defaultModel

```go
const defaultModel = "gpt-4o-mini"
```

The default OpenAI model used when no model is specified via `--model`.

### apiURL

```go
const apiURL = "https://api.openai.com/v1/chat/completions"
```

The OpenAI API endpoint for chat completions.

## Usage Examples

### Basic Usage

```bash
# Simple processing
echo "Hello, world!" | sai

# With context
echo "Paris" | sai --context "List 3 famous landmarks in"
```

### Advanced Usage

```bash
# Use different model
echo "What is machine learning?" | sai --model gpt-3.5-turbo

# Code generation (clean output)
echo "Python function to calculate fibonacci" | sai --code

# Custom API key
echo "Translate to Spanish: Hello" | sai --api-key sk-your-key
```

### File Processing

```bash
# Process file content
cat document.txt | sai --context "Summarize this document:"

# Code review
cat main.go | sai --context "Review this Go code and suggest improvements:"

# Data transformation
cat data.json | sai --context "Convert this JSON to YAML format:"
```

### Chaining Operations

```bash
# Multi-step processing
echo "Create a REST API" | sai --code | sai --context "Add error handling to this code:"

# Documentation generation
cat *.go | sai --context "Generate API documentation for this Go code:"
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

3. **API Error Response**
   ```
   Error from API (status 401): Invalid API key
   ```

4. **Network/Connection Issues**
   ```
   Error sending request: context deadline exceeded
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

# Handle potential errors
echo "test input" | sai || echo "Failed to process input"
```

## Development Guide

### Building

```bash
# Build for current platform
make build

# Build for all platforms (requires cross-compilation setup)
make build-all

# Clean build artifacts
make clean
```

### Testing

```bash
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
│       └── release.yml          # CI/CD pipeline
├── Makefile                     # Build automation
├── go.mod                       # Go module definition
└── README.md                    # Project documentation
```

### Adding New Features

1. **Command-line flags**: Add new flags in the flag parsing section
2. **Request modification**: Extend the `Request` type and modify request construction
3. **Response processing**: Update the streaming response handling logic
4. **Error handling**: Add specific error cases and user-friendly messages

### Release Process

1. Tag the commit:
   ```bash
   git tag -a v1.0.0 -m "Release v1.0.0"
   git push origin v1.0.0
   ```

2. GitHub Actions automatically builds and releases binaries for all platforms

### Dependencies

- **Standard Library Only**: The project uses only Go's standard library
- **No External Dependencies**: Ensures easy compilation and distribution
- **HTTP Client**: Uses `net/http` for OpenAI API communication
- **JSON Processing**: Uses `encoding/json` for request/response handling

### Performance Considerations

- **Streaming**: Real-time response streaming for immediate feedback
- **Memory Efficient**: Minimal memory footprint with streaming processing
- **Single Binary**: No runtime dependencies required
- **Cross-platform**: Native compilation for all major platforms

## API Integration Details

### OpenAI API Usage

The tool integrates with OpenAI's Chat Completions API using:

- **Endpoint**: `https://api.openai.com/v1/chat/completions`
- **Method**: POST
- **Authentication**: Bearer token in Authorization header
- **Streaming**: Server-sent events (SSE) for real-time responses

### Request Format

```json
{
  "model": "gpt-4o-mini",
  "messages": [
    {
      "role": "user",
      "content": "Your input text here"
    }
  ],
  "stream": true
}
```

### Response Format

```json
{
  "choices": [
    {
      "delta": {
        "content": "Incremental text content"
      }
    }
  ]
}
```

---

**Version**: 1.0.0  
**License**: MIT  
**Author**: rwejlgaard  
**Repository**: https://github.com/rwejlgaard/sai
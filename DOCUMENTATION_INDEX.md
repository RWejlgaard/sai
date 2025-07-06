# sai Documentation Index

This document provides a comprehensive overview of all documentation for the `sai` (Stream AI) project.

## 📚 Documentation Overview

The sai project documentation is organized into several focused documents:

### 1. **API_DOCUMENTATION.md** - User-Focused API Documentation
- **Purpose**: Complete user guide and API reference
- **Audience**: End users, system administrators, DevOps engineers
- **Contents**:
  - Installation instructions
  - Command-line API reference
  - Usage examples and patterns
  - Error handling guide
  - Configuration and environment variables

### 2. **DEVELOPER_API.md** - Developer-Focused Technical Reference
- **Purpose**: Technical documentation for developers
- **Audience**: Go developers, contributors, maintainers
- **Contents**:
  - Code architecture and structure
  - Type definitions with examples
  - Function-level documentation
  - Extension points and customization
  - Performance characteristics

### 3. **README.md** - Project Overview
- **Purpose**: Project introduction and quick start
- **Audience**: General users, new contributors
- **Contents**:
  - Project description
  - Basic installation and usage
  - Development workflow
  - Release process

## 🚀 Quick Start

### For End Users
```bash
# Install
curl -L https://github.com/rwejlgaard/sai/releases/latest/download/sai-linux-amd64 -o sai
chmod +x sai && sudo mv sai /usr/local/bin/

# Configure
export OPENAI_API_KEY='your-api-key-here'

# Use
echo "Hello, world!" | sai
```

### For Developers
```bash
# Clone and build
git clone https://github.com/rwejlgaard/sai.git
cd sai
make build

# Run tests
make test

# See DEVELOPER_API.md for detailed technical documentation
```

## 📖 Documentation Structure

### Public APIs and Interfaces

#### Command-Line Interface
- **Binary**: `sai`
- **Input**: stdin (pipe or redirect)
- **Output**: stdout (streaming)
- **Flags**: `--context`, `--model`, `--api-key`, `--code`
- **Environment**: `OPENAI_API_KEY`

#### Data Types (Go)
- `Message`: OpenAI chat message structure
- `Request`: API request payload
- `Response`: API response structure
- `Delta`: Incremental content for streaming
- `Choice`: Response choice container

#### Constants
- `defaultModel`: `"gpt-4o-mini"`
- `apiURL`: `"https://api.openai.com/v1/chat/completions"`

### Internal Architecture

#### Core Components
1. **Flag Parser**: Command-line argument processing
2. **Input Processor**: Stdin reading and validation
3. **HTTP Client**: OpenAI API integration
4. **Stream Handler**: Real-time response processing

#### Key Functions
- `main()`: Application entry point
- Flag parsing functions
- HTTP request/response handling
- Streaming response processing

## 🔧 Configuration Options

### Command-Line Flags

| Flag | Short | Type | Default | Description |
|------|-------|------|---------|-------------|
| `--context` | `-c` | string | `""` | Context to prepend to input |
| `--model` | `-m` | string | `"gpt-4o-mini"` | OpenAI model to use |
| `--api-key` | `-k` | string | `""` | OpenAI API key |
| `--code` | `-o` | bool | `false` | Return only code |

### Environment Variables

| Variable | Required | Description |
|----------|----------|-------------|
| `OPENAI_API_KEY` | Yes* | OpenAI API key for authentication |

*Required unless `--api-key` flag is provided

## 📝 Usage Examples

### Basic Usage
```bash
# Simple AI processing
echo "Explain quantum computing" | sai

# With context
echo "Python" | sai -c "Write a hello world program in"

# Code-only output
echo "fibonacci function" | sai --code
```

### Advanced Usage
```bash
# File processing
cat document.txt | sai -c "Summarize this document:"

# Different model
echo "Complex math problem" | sai -m gpt-4

# Chained processing
echo "REST API" | sai --code | sai -c "Add error handling:"
```

## 🏗️ Architecture Overview

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Input (stdin) │───▶│  sai CLI Tool   │───▶│ Output (stdout) │
└─────────────────┘    └─────────────────┘    └─────────────────┘
                              │
                              ▼
                    ┌─────────────────┐
                    │   OpenAI API    │
                    │  (Streaming)    │
                    └─────────────────┘
```

### Data Flow
1. **Input**: Read from stdin (pipe/redirect)
2. **Processing**: Add context, prepare request
3. **API Call**: Send to OpenAI with streaming enabled
4. **Response**: Stream output to stdout in real-time

## 🔍 Error Handling

### Common Errors
- Missing API key
- No input provided
- Invalid API key
- Network connectivity issues
- API rate limits

### Error Codes
- `0`: Success
- `1`: Configuration or runtime error

### Best Practices
- Set `OPENAI_API_KEY` environment variable
- Use proper input validation
- Handle network timeouts gracefully
- Monitor API usage and costs

## 🛠️ Development

### Build Requirements
- Go 1.22.5+
- No external dependencies (standard library only)

### Build Commands
```bash
make build    # Build binary
make test     # Run tests
make clean    # Clean artifacts
make install  # Install system-wide
```

### Project Structure
```
sai/
├── src/cmd/sai/main.go     # Main application
├── .github/workflows/      # CI/CD automation
├── Makefile               # Build automation
├── go.mod                 # Go module
└── *.md                   # Documentation
```

## 📊 Performance Characteristics

### Memory Usage
- **Minimal**: Streaming processing with low memory footprint
- **Efficient**: Uses `strings.Builder` for string concatenation
- **No Caching**: Stateless operation per request

### Network Efficiency
- **Streaming**: Real-time response output
- **Single Request**: One HTTP connection per invocation
- **Connection Reuse**: Default HTTP client pooling

### Binary Size
- **Small**: ~2-5MB static binary
- **Zero Dependencies**: No runtime requirements
- **Cross-Platform**: Native compilation for all platforms

## 🔗 Links and References

### Documentation Files
- [API_DOCUMENTATION.md](./API_DOCUMENTATION.md) - Complete user guide
- [DEVELOPER_API.md](./DEVELOPER_API.md) - Technical developer reference
- [README.md](./README.md) - Project overview and quick start

### External Resources
- [OpenAI API Documentation](https://platform.openai.com/docs/api-reference)
- [Go Documentation](https://golang.org/doc/)
- [GitHub Repository](https://github.com/rwejlgaard/sai)

### Related Tools
- `curl` - HTTP client for API testing
- `jq` - JSON processing for API responses
- `sed` - Stream editor (inspiration for sai)

---

**Last Updated**: Generated automatically  
**Version**: 1.0.0  
**Maintainer**: rwejlgaard
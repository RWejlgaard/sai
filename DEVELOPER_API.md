# sai Developer API Reference

This document provides detailed technical documentation for developers working with the `sai` codebase.

## Package Overview

```go
package main
```

The `sai` application is implemented as a single-package Go program located at `src/cmd/sai/main.go`.

## Imports

The application uses only Go standard library packages:

```go
import (
    "bufio"       // Buffered I/O operations
    "bytes"       // Byte buffer operations
    "encoding/json" // JSON encoding/decoding
    "flag"        // Command-line flag parsing
    "fmt"         // Formatted I/O
    "io"          // I/O primitives
    "net/http"    // HTTP client
    "os"          // Operating system interface
    "strings"     // String manipulation
)
```

## Constants

### defaultModel

```go
const defaultModel = "gpt-4o-mini"
```

- **Type**: `string`
- **Purpose**: Default OpenAI model identifier used when no model is specified
- **Usage**: Referenced in flag parsing and request construction

### apiURL

```go
const apiURL = "https://api.openai.com/v1/chat/completions"
```

- **Type**: `string`
- **Purpose**: OpenAI API endpoint for chat completions
- **Usage**: Used in HTTP request construction

## Type Definitions

### Message

```go
type Message struct {
    Role    string `json:"role"`
    Content string `json:"content"`
}
```

**Purpose**: Represents a single message in the OpenAI chat format

**Fields**:
- `Role`: Message role (`"user"`, `"assistant"`, or `"system"`)
- `Content`: The actual message content

**JSON Tags**: Properly tagged for JSON serialization

**Usage Example**:
```go
userMessage := Message{
    Role:    "user",
    Content: "Hello, world!",
}
```

### Request

```go
type Request struct {
    Model    string    `json:"model"`
    Messages []Message `json:"messages"`
    Stream   bool      `json:"stream"`
}
```

**Purpose**: Represents the complete request payload sent to OpenAI's API

**Fields**:
- `Model`: OpenAI model identifier (e.g., `"gpt-4o-mini"`)
- `Messages`: Array of conversation messages
- `Stream`: Always set to `true` for streaming responses

**Usage Example**:
```go
req := Request{
    Model: "gpt-4o-mini",
    Messages: []Message{
        {Role: "user", Content: input},
    },
    Stream: true,
}
```

### Delta

```go
type Delta struct {
    Content string `json:"content"`
}
```

**Purpose**: Represents incremental content in streaming responses

**Fields**:
- `Content`: Incremental text content from the API

### Choice

```go
type Choice struct {
    Delta Delta `json:"delta"`
}
```

**Purpose**: Represents a single choice/option in the API response

**Fields**:
- `Delta`: Contains incremental content for streaming responses

### Response

```go
type Response struct {
    Choices []Choice `json:"choices"`
}
```

**Purpose**: Represents the complete response structure from OpenAI's API

**Fields**:
- `Choices`: Array of response choices (typically contains one choice)

## Main Function

### Function Signature

```go
func main()
```

**Purpose**: Entry point of the application

**Flow**:
1. Parse command-line flags
2. Validate API key
3. Check for stdin input
4. Read input from stdin
5. Construct prompt with context
6. Create and send HTTP request
7. Process streaming response
8. Output results to stdout

## Command-Line Flag Parsing

The application uses Go's `flag` package for command-line argument processing:

```go
context := flag.String("context", "", "Context to prepend to the input (shorthand: -c)")
model := flag.String("model", defaultModel, "OpenAI model to use (shorthand: -m)")
apiKey := flag.String("api-key", "", "OpenAI API key (shorthand: -k)")
codeOnly := flag.Bool("code", false, "Return only code without explanations or backticks (shorthand: -o)")

// Short flag variants
flag.StringVar(context, "c", "", "Context to prepend to the input")
flag.StringVar(model, "m", defaultModel, "OpenAI model to use")
flag.StringVar(apiKey, "k", "", "OpenAI API key")
flag.BoolVar(codeOnly, "o", false, "Return only code without explanations or backticks")
```

**Flag Types**:
- `string`: For text-based options
- `bool`: For boolean switches

## Input Processing

### Stdin Validation

```go
stat, _ := os.Stdin.Stat()
if (stat.Mode() & os.ModeCharDevice) != 0 {
    // Error: no piped input
}
```

**Purpose**: Checks if input is available from stdin (pipe or redirect)

**Logic**: Uses `os.Stdin.Stat()` to detect if stdin is a character device (terminal) or has piped input

### Input Reading

```go
reader := bufio.NewReader(os.Stdin)
var input strings.Builder
for {
    line, err := reader.ReadString('\n')
    if err == io.EOF {
        break
    }
    if err != nil {
        // Handle error
    }
    input.WriteString(line)
}
```

**Purpose**: Reads all input from stdin line by line

**Components**:
- `bufio.NewReader()`: Creates buffered reader for efficient I/O
- `strings.Builder`: Efficient string concatenation
- `ReadString('\n')`: Reads until newline character

## API Integration

### HTTP Request Construction

```go
httpReq, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(reqBody))
if err != nil {
    // Handle error
}

httpReq.Header.Set("Content-Type", "application/json")
httpReq.Header.Set("Authorization", "Bearer "+*apiKey)
```

**Purpose**: Creates HTTP POST request to OpenAI API

**Headers**:
- `Content-Type`: Set to `application/json`
- `Authorization`: Bearer token authentication

### Request Body Marshaling

```go
reqBody, err := json.Marshal(req)
if err != nil {
    // Handle error
}
```

**Purpose**: Converts `Request` struct to JSON bytes

**Error Handling**: Exits with error code 1 if marshaling fails

### HTTP Client Usage

```go
client := &http.Client{}
resp, err := client.Do(httpReq)
if err != nil {
    // Handle error
}
defer resp.Body.Close()
```

**Purpose**: Sends HTTP request and handles response

**Cleanup**: Uses `defer` to ensure response body is closed

## Response Processing

### Status Code Validation

```go
if resp.StatusCode != http.StatusOK {
    body, _ := io.ReadAll(resp.Body)
    fmt.Fprintf(os.Stderr, "Error from API (status %d): %s\n", resp.StatusCode, string(body))
    os.Exit(1)
}
```

**Purpose**: Validates API response status and handles errors

**Error Output**: Prints error details to stderr

### Streaming Response Handling

```go
scanner := bufio.NewScanner(resp.Body)
for scanner.Scan() {
    line := scanner.Text()
    if !strings.HasPrefix(line, "data: ") {
        continue
    }
    
    data := strings.TrimPrefix(line, "data: ")
    if data == "" || data == "[DONE]" {
        continue
    }
    
    var response Response
    if err := json.Unmarshal([]byte(data), &response); err != nil {
        // Handle error
        continue
    }
    
    if len(response.Choices) > 0 && response.Choices[0].Delta.Content != "" {
        fmt.Print(response.Choices[0].Delta.Content)
    }
}
```

**Purpose**: Processes server-sent events (SSE) from OpenAI API

**Flow**:
1. Scan response body line by line
2. Filter for lines starting with "data: "
3. Skip empty lines and "[DONE]" markers
4. Unmarshal JSON to Response struct
5. Extract and print content from first choice

## Error Handling Patterns

### API Key Validation

```go
if *apiKey == "" {
    *apiKey = os.Getenv("OPENAI_API_KEY")
    if *apiKey == "" {
        fmt.Fprintln(os.Stderr, "Error: OpenAI API key must be provided either via --api-key or OPENAI_API_KEY environment variable")
        os.Exit(1)
    }
}
```

**Strategy**: Fallback from flag to environment variable

### Input Validation

```go
stat, _ := os.Stdin.Stat()
if (stat.Mode() & os.ModeCharDevice) != 0 {
    fmt.Fprintln(os.Stderr, "Error: No input provided. Please pipe some input to sai.")
    fmt.Fprintln(os.Stderr, "Example: echo 'Hello, world!' | sai")
    os.Exit(1)
}
```

**Strategy**: Provide helpful error messages with usage examples

### Network Error Handling

```go
if err := scanner.Err(); err != nil {
    fmt.Fprintf(os.Stderr, "Error reading response: %v\n", err)
    os.Exit(1)
}
```

**Strategy**: Check for scanner errors after processing

## Context Processing

### Basic Context

```go
prompt := input.String()
if *context != "" {
    prompt = *context + "\n\n" + prompt
}
```

**Purpose**: Prepends user-provided context to input

**Format**: Context and input separated by double newlines

### Code-Only Mode

```go
if *codeOnly {
    prompt = "You are a code-only assistant. Return ONLY the code without any explanations, backticks, or markdown formatting. Do not include any text before or after the code.\n\n" + prompt
}
```

**Purpose**: Modifies prompt to request code-only responses

**Strategy**: System-level instruction prepended to user input

## Performance Characteristics

### Memory Usage

- **Streaming**: Processes response incrementally, minimal memory footprint
- **Input Buffering**: Uses `strings.Builder` for efficient string concatenation
- **No Caching**: Each request is independent, no persistent state

### Network Efficiency

- **Streaming**: Real-time output as content arrives
- **Single Connection**: One HTTP request per invocation
- **Connection Reuse**: Default HTTP client connection pooling

### Error Recovery

- **Graceful Degradation**: Continues processing even if individual chunks fail
- **Clear Error Messages**: User-friendly error reporting
- **Fast Failure**: Exits quickly on configuration errors

## Extension Points

### Adding New Flags

```go
// Add new flag
newFlag := flag.String("new-flag", "default", "Description")
flag.StringVar(newFlag, "n", "default", "Description")

// Use in request construction
if *newFlag != "default" {
    // Modify request or processing
}
```

### Modifying Request Structure

```go
// Extend Request type
type Request struct {
    Model       string    `json:"model"`
    Messages    []Message `json:"messages"`
    Stream      bool      `json:"stream"`
    Temperature float64   `json:"temperature,omitempty"` // New field
}
```

### Adding Response Processing

```go
// In streaming loop
if len(response.Choices) > 0 {
    choice := response.Choices[0]
    if choice.Delta.Content != "" {
        // Custom processing here
        processContent(choice.Delta.Content)
    }
}
```

## Build and Deployment

### Build Process

```bash
# Standard build
go build -o dist/sai ./src/cmd/sai

# Cross-compilation example
GOOS=linux GOARCH=amd64 go build -o dist/sai-linux-amd64 ./src/cmd/sai
```

### Binary Size Optimization

The application is optimized for small binary size:
- **No External Dependencies**: Uses only standard library
- **Single File**: All code in one source file
- **Minimal Imports**: Only necessary packages imported

### Distribution

- **Static Binary**: No runtime dependencies
- **Cross-Platform**: Builds for all major platforms
- **GitHub Releases**: Automated binary distribution

---

**File Location**: `src/cmd/sai/main.go`  
**Package**: `main`  
**Go Version**: 1.22.5+  
**Dependencies**: Go standard library only
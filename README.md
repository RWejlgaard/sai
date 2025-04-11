# sai (Stream AI)

A command-line utility that processes input streams through OpenAI's API, similar to `sed` but for AI processing. Written in Go for maximum performance and easy distribution.

## Installation

1. Clone this repository
2. Build and install:
   ```bash
   make
   sudo make install
   ```

Or install directly using Go:
```bash
go install github.com/yourusername/sai/src/cmd/sai@latest
```

## Usage

Set your OpenAI API key as an environment variable:
```bash
export OPENAI_API_KEY='your-api-key-here'
```

Basic usage:
```bash
echo "Hello, world!" | sai
```

With context:
```bash
echo "Paris" | sai --context "List 3 famous landmarks in"
# or using the short flag
echo "Paris" | sai -c "List 3 famous landmarks in"
```

Using a different model:
```bash
echo "What is 2+2?" | sai --model gpt-3.5-turbo
# or using the short flag
echo "What is 2+2?" | sai -m gpt-3.5-turbo
```

Using a different API key:
```bash
echo "Hello" | sai --api-key 'your-alternative-api-key'
# or using the short flag
echo "Hello" | sai -k 'your-alternative-api-key'
```

## Arguments

- `--context`, `-c`: Optional context to prepend to the input
- `--model`, `-m`: OpenAI model to use (default: gpt-4-turbo-preview)
- `--api-key`, `-k`: OpenAI API key (defaults to OPENAI_API_KEY environment variable)

## Examples

1. Summarize a file:
   ```bash
   cat document.txt | sai --context "Summarize the following text:"
   ```

2. Generate code from description:
   ```bash
   echo "A Go function that calculates fibonacci numbers" | sai
   ```

3. Translate text:
   ```bash
   echo "Hello, how are you?" | sai --context "Translate to Spanish:"
   ```

## Development

- Build: `make build`
- Install: `sudo make install`
- Clean: `make clean`
- Test: `make test`
- Format code: `make format`
- Lint code: `make lint` 
# sai (Stream AI)

A command-line utility that processes input streams through OpenAI's API, similar to `sed` but for AI processing. Written in Go for maximum performance and easy distribution.

## Installation

### Pre-built binaries (recommended)

Download the latest release for your platform from the [releases page](https://github.com/yourusername/sai/releases/latest):

```bash
# For macOS (Intel)
curl -L https://github.com/rwejlgaard/sai/releases/latest/download/sai-darwin-amd64 -o sai
chmod +x sai
sudo mv sai /usr/local/bin/

# For macOS (Apple Silicon)
curl -L https://github.com/rwejlgaard/sai/releases/latest/download/sai-darwin-arm64 -o sai
chmod +x sai
sudo mv sai /usr/local/bin/

# For Linux (x86_64)
curl -L https://github.com/rwejlgaard/sai/releases/latest/download/sai-linux-amd64 -o sai
chmod +x sai
sudo mv sai /usr/local/bin/

# For Linux (ARM64)
curl -L https://github.com/rwejlgaard/sai/releases/latest/download/sai-linux-arm64 -o sai
chmod +x sai
sudo mv sai /usr/local/bin/
```

For Windows users, download the appropriate `.exe` file from the releases page and add it to your PATH.

### Building from source

1. Clone this repository
2. Build and install:
   ```bash
   make
   sudo make install
   ```

Or install directly using Go:
```bash
go install github.com/rwejlgaard/sai/src/cmd/sai@latest
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

## Releasing

To create a new release:

1. Tag the commit:
   ```bash
   git tag -a v1.0.0 -m "Release v1.0.0"
   ```

2. Push the tag:
   ```bash
   git push origin v1.0.0
   ```

This will trigger the GitHub Actions workflow to build and release the binaries. 
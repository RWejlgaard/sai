package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	defaultModel = "gpt-4o-mini"
	apiURL       = "https://api.openai.com/v1/chat/completions"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Request struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
	Stream   bool      `json:"stream"`
}

type Delta struct {
	Content string `json:"content"`
}

type Choice struct {
	Delta Delta `json:"delta"`
}

type Response struct {
	Choices []Choice `json:"choices"`
}

func main() {
	// Parse command line flags
	context := flag.String("context", "", "Context to prepend to the input (shorthand: -c)")
	model := flag.String("model", defaultModel, "OpenAI model to use (shorthand: -m)")
	apiKey := flag.String("api-key", "", "OpenAI API key (shorthand: -k)")
	codeOnly := flag.Bool("code", false, "Return only code without explanations or backticks (shorthand: -o)")
	flag.StringVar(context, "c", "", "Context to prepend to the input")
	flag.StringVar(model, "m", defaultModel, "OpenAI model to use")
	flag.StringVar(apiKey, "k", "", "OpenAI API key")
	flag.BoolVar(codeOnly, "o", false, "Return only code without explanations or backticks")
	flag.Parse()

	// Get API key from environment if not provided
	if *apiKey == "" {
		*apiKey = os.Getenv("OPENAI_API_KEY")
		if *apiKey == "" {
			fmt.Fprintln(os.Stderr, "Error: OpenAI API key must be provided either via --api-key or OPENAI_API_KEY environment variable")
			os.Exit(1)
		}
	}

	// Check if stdin has data available
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		fmt.Fprintln(os.Stderr, "Error: No input provided. Please pipe some input to sai.")
		fmt.Fprintln(os.Stderr, "Example: echo 'Hello, world!' | sai")
		os.Exit(1)
	}

	// Read from stdin
	reader := bufio.NewReader(os.Stdin)
	var input strings.Builder
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
			os.Exit(1)
		}
		input.WriteString(line)
	}

	// Construct prompt
	prompt := input.String()
	if *context != "" {
		prompt = *context + "\n\n" + prompt
	}
	if *codeOnly {
		prompt = "You are a code-only assistant. Return ONLY the code without any explanations, backticks, or markdown formatting. Do not include any text before or after the code.\n\n" + prompt
	}

	// Create request
	req := Request{
		Model: *model,
		Messages: []Message{
			{
				Role:    "user",
				Content: prompt,
			},
		},
		Stream: true,
	}

	// Marshal request to JSON
	reqBody, err := json.Marshal(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling request: %v\n", err)
		os.Exit(1)
	}

	// Create HTTP request
	httpReq, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating request: %v\n", err)
		os.Exit(1)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+*apiKey)

	// Send request
	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error sending request: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		fmt.Fprintf(os.Stderr, "Error from API (status %d): %s\n", resp.StatusCode, string(body))
		os.Exit(1)
	}

	// Process streaming response
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "data: ") {
			continue
		}

		// Skip empty lines and "[DONE]" messages
		data := strings.TrimPrefix(line, "data: ")
		if data == "" || data == "[DONE]" {
			continue
		}

		var response Response
		if err := json.Unmarshal([]byte(data), &response); err != nil {
			fmt.Fprintf(os.Stderr, "Error decoding response: %v\n", err)
			continue
		}

		if len(response.Choices) > 0 && response.Choices[0].Delta.Content != "" {
			fmt.Print(response.Choices[0].Delta.Content)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading response: %v\n", err)
		os.Exit(1)
	}

	fmt.Println() // Add final newline
}

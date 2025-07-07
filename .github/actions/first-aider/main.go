package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
)

// ErrorMapping represents a description and fix for a known error.
type ErrorMapping struct {
	Description string `json:"description"`
	Fix         string `json:"fix"`
}

type ClaudeRequest struct {
	Model     string      `json:"model"`
	Messages  []ClaudeMsg `json:"messages"`
	MaxTokens int         `json:"max_tokens"`
	Stream    bool        `json:"stream"`
}

type ClaudeMsg struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func main() {
	ctx := context.Background()

	// Parse input flag
	logPath := flag.String("log-path", "", "Path to the build log file")
	flag.Parse()

	if *logPath == "" {
		fmt.Fprintln(os.Stderr, "Error: --log-path is required")
		os.Exit(1)
	}

	// Read log file
	rawLog, err := os.ReadFile(*logPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read log file: %v\n", err)
		os.Exit(1)
	}

	// Extract relevant snippet: last 100 lines
	lines := strings.Split(string(rawLog), "\n")
	snippet := ""
	if len(lines) > 100 {
		snippet = strings.Join(lines[len(lines)-100:], "\n")
	} else {
		snippet = string(rawLog)
	}

	// Load error mappings
	mappingData, err := os.ReadFile("errors.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read errors.json: %v\n", err)
		os.Exit(1)
	}
	var allMappings map[string]ErrorMapping
	if err := json.Unmarshal(mappingData, &allMappings); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse errors.json: %v\n", err)
		os.Exit(1)
	}

	// Filter mappings relevant to the snippet
	relevant := make(map[string]ErrorMapping)
	for key, info := range allMappings {
		if strings.Contains(snippet, key) {
			relevant[key] = info
		}
	}
	relevantJSON, _ := json.MarshalIndent(relevant, "", "  ")

	// Compose Claude prompt
	prompt := fmt.Sprintf(
		"You are a CI First-Aider. A Go CI job failed with the following log snippet:\n```\n%s\n```\n", snippet) +
		fmt.Sprintf("Known error mappings from errors.json (if key appears in snippet):\n```json\n%s\n```\n", relevantJSON) +
		"Analyze the root cause of the failure and suggest 2-3 actionable fixes in Markdown format. Provide one code snippet max." +
		"The first data should be a table for the user. It should be 3 columns: file, line the error appeared on and cause." +
		"Each failure should have a row in the table. If the logs show no errors, write a congratulation message to the dev."

	// Delegate to Connect and capture suggestion
	suggestion, err := connect(ctx, prompt)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting suggestion: %v\n", err)
		os.Exit(1)
	}

	for _, s := range suggestion {
		fmt.Printf("Resp: %s", s.Text)
	}

	// Emit output for GitHub Action
	// fmt.Printf("::set-output name=comment::%+v", suggestion)
}

func connect(ctx context.Context, prompt string) ([]anthropic.ContentBlockUnion, error) {
	flag.Parse()

	// Call Claude via Anthropic SDK
	key := os.Getenv("CLAUDE_TOKEN")
	client := anthropic.NewClient(option.WithAPIKey(key))
	respMsg, err := client.Messages.New(ctx, anthropic.MessageNewParams{
		Model:     anthropic.ModelClaude4Opus20250514,
		MaxTokens: 500,
		Messages: []anthropic.MessageParam{
			anthropic.NewUserMessage(anthropic.NewTextBlock(prompt)),
		},
	})
	if err != nil {
		return nil, fmt.Errorf("Error calling Claude API: %w", err)
	}
	suggestion := respMsg.Content
	return suggestion, nil
}

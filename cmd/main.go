package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"napier/hack/ai"
	"napier/hack/errors"
	"os"
	"path/filepath"
	"strings"
)

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
	snippet := getErrLines(string(rawLog))

	// Load error mappings
	mapPath := getErrMap()

	mappingData, err := os.ReadFile(mapPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read %s: %v\n", mapPath, err)
		os.Exit(1)
	}
	var allMappings map[string]errors.ErrorMapping
	if err := json.Unmarshal(mappingData, &allMappings); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse errors.json: %v\n", err)
		os.Exit(1)
	}

	// Filter mappings relevant to the snippet
	relevant := make(map[string]errors.ErrorMapping)
	for key, info := range allMappings {
		if strings.Contains(snippet, key) {
			relevant[key] = info
		}
	}
	relevantJSON, _ := json.MarshalIndent(relevant, "", "  ")

	// Compose Claude prompt
	prompt := ai.BuildPrompt(snippet, relevantJSON)

	// Delegate to Connect and capture suggestion
	suggestions, err := ai.Connect(ctx, prompt)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting suggestion: %v\n", err)
		os.Exit(1)
	}

	var suggestedAction string
	for _, suggestion := range suggestions {
		suggestedAction = suggestion.Text
	}

	// Emit output for GitHub Action
	fmt.Printf("::set-output name=comment::%s", suggestedAction)
}

// getErrLines returns the last 100 lines of an error log
func getErrLines(rawLog string) string {
	lines := strings.Split(rawLog, "\n")
	snippet := ""
	if len(lines) > 100 {
		snippet = strings.Join(lines[len(lines)-100:], "\n")
	} else {
		snippet = rawLog
	}
	return snippet
}

// getErrMap returns the path to the errors map
func getErrMap() string {
	exePath, err := os.Executable()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to find executable path: %v\n", err)
		os.Exit(1)
	}
	exeDir := filepath.Dir(exePath)

	return filepath.Join(exeDir, "errors/errors.json")
}

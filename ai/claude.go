package ai

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
)

func Connect(ctx context.Context, prompt string) ([]anthropic.ContentBlockUnion, error) {
	flag.Parse()

	// Call Claude via Anthropic SDK
	client := anthropic.NewClient(option.WithAPIKey(os.Getenv("CLAUDE_KEY")))
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

func BuildPrompt(snippet string, errJSON []byte) string {
	// limitation items
	title := "A title indicating the severity of the fail (pass, low, medium, high)"
	table := "A table with 3 columns (the file name, line the error appeared on and the cause) outlining each error on it own row"
	passMsg := "If the log shows no errors, write a one line congratulations message for the developer."

	// basic prompt
	taskCtx := "Your role is to assist developers in quickly debugging a failed pipeline stage - you are to act as a pipeline first aider."
	errLog := fmt.Sprintf("A CI job/stage has failed and here is the last 100 lines of the log message: \n```\n%s\n```\n", snippet)
	errMap := fmt.Sprintf("Known error mappings from errors.json (if key appears in snippet):\n```json\n%s\n```\n", errJSON)
	instruction := "You are to analyze the root cause of the failure and suggest an actionable fix in Markdown format."
	limitation := fmt.Sprintf("Your response should use the following template: %s, %s, %s", title, table, passMsg)

	return taskCtx + errLog + errMap + instruction + limitation
}

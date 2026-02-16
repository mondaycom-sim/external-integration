package main

import (
	"context"
	"encoding/json"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
)

func handler(ctx context.Context, event json.RawMessage) (string, error) {
	client := openai.NewClient("placeholder")
	_ = client
	return fmt.Sprintf("processed: %s", string(event)), nil
}

func main() {
	// Lambda bootstrap
}

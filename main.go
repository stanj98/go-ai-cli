package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/openai/openai-go/v2"
	"github.com/openai/openai-go/v2/option"
)

func main() {
	apiKey := os.Getenv("OPENROUTER_API_KEY")
	if apiKey == "" {
		log.Fatal("API key is required")
	}
	url := "https://openrouter.ai/api/v1"
	client := openai.NewClient(
		option.WithBaseURL(url),
		option.WithAPIKey(apiKey),
	)

	ctx := context.Background()

	//slice of openai chat completion messages
	messages := []openai.ChatCompletionMessageParamUnion{}

	model := "x-ai/grok-4-fast:free"

	//append query
	messages = append(messages, openai.UserMessage("What model are you?"))

	params := openai.ChatCompletionNewParams{
		Model:    model,
		Messages: messages,
	}

	res, err := client.Chat.Completions.New(ctx, params)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(res.Choices[0].Message.Content)

}

package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

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

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}
		if input == "clear" {
			fmt.Print("\033[H\033[2J")
			continue
		}
		if input == "exit" {
			fmt.Println("goodbye")
			break
		}

		params.Messages = append(params.Messages, openai.UserMessage(input))

		res, err := client.Chat.Completions.New(ctx, params)
		if err != nil {
			log.Fatal(err)
		}

		//append the answer of LLM to params messages
		output := res.Choices[0].Message.Content
		fmt.Printf(output)

		params.Messages = append(params.Messages, openai.AssistantMessage(output))
	}

}

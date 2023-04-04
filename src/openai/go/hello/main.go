package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
)

func request(ctx context.Context, model string, messages []openai.ChatCompletionMessage) openai.ChatCompletionRequest {
	return openai.ChatCompletionRequest{
		Model:     model,
		MaxTokens: 1000,
		Messages:  messages,
		Stream:    true,
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	OPENAI_API_KEY := os.Getenv("OPENAI_API_KEY")

	client := openai.NewClient(OPENAI_API_KEY)
	ctx := context.Background()
	model := openai.GPT3Dot5Turbo

	messages := make([]openai.ChatCompletionMessage, 0)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the OpenAI Chatbot!")
	fmt.Println("Type 'exit' to quit.")
	fmt.Println()

	for {
		fmt.Print("You: ")
		text, _ := reader.ReadString('\n')
		text = text[:len(text)-1]

		if text == "exit" {
			break
		}

		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: text,
		})

		req := request(ctx, model, messages)
		stream, err := client.CreateChatCompletionStream(ctx, req)
		if err != nil {
			fmt.Printf("Chat Stream Error: %v\n", err)
			continue
		}
		defer stream.Close()

		fmt.Printf("Bot: ")

		for {
			resp, err := stream.Recv()
			if errors.Is(err, io.EOF) {
				fmt.Println()
				break
			}
			if err != nil {
				fmt.Printf("\nStream error: %v\n", err)
				return
			}

			content := resp.Choices[0].Delta.Content

			messages = append(messages, openai.ChatCompletionMessage{
				Role:    openai.ChatMessageRoleAssistant,
				Content: content,
			})

			fmt.Printf("%s", content)
		}
		fmt.Println()

	}
}

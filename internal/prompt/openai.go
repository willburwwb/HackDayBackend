package prompt

import (
	"context"
	"fmt"
	openai "github.com/sashabaranov/go-openai"
	"os"
)

func getMessageFromAI(content string) (string, error) {
	client := openai.NewClient(os.Getenv("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"))
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: content,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "", err
	}

	fmt.Println(resp.Choices[0].Message.Content)
	return resp.Choices[0].Message.Content, nil
}

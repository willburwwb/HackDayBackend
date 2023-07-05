package prompt

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"log"
	"os"
)

func getMessageFromAI(content string) (string, error) {
	client := openai.NewClient(os.Getenv("OPENAI_TOKEN"))
	log.Println(os.Getenv("OPENAI_TOKEN"), content)
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

	//Todo
	// 由于在docker里莫名其妙跑不通，在服务器增加一个http服务openai接口，这里通过http client调用服务，
	//return "test1", nil
}

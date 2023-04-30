// Package provides a Client for GPT API.
package gpt

import (
	"context"
	"os"

	"github.com/sashabaranov/go-openai"
)

// Represents a client for GPT API.
type Client struct {
	Client *openai.Client
}

// A default client for GPT API.
var DefaultClient = NewClient(os.Getenv("SCHEDULER_GPT_API_TOKEN"))

// Returns a new client for GPT API
func NewClient(apiKey string) Client {
	return Client{Client: openai.NewClient(apiKey)}
}

// Returns a new client for GPT API with specified API url.
func NewClientWithURL(apiKey, url string) Client {
	config := openai.DefaultConfig(apiKey)
	config.BaseURL = url

	return Client{Client: openai.NewClientWithConfig(config)}
}

// Returns prompt completion.
func Complete(prompt string) (string, error) {
	return DefaultClient.Complete(prompt)
}

// Returns prompt completion.
func (c Client) Complete(prompt string) (string, error) {
	resp, err := c.Client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}

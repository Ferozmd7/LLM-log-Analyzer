package controllers

import "fmt"

// LLMClient defines a common interface for any LLM
type LLMClient interface {
	Generate(prompt string) (string, error)
}

// -------------------
// OpenAI Implementation
// -------------------
type OpenAIClient struct {
	APIKey string
}

func NewOpenAIClient(apiKey string) *OpenAIClient {
	return &OpenAIClient{APIKey: apiKey}
}

func (c *OpenAIClient) Generate(prompt string) (string, error) {
	// Placeholder: implement API call to OpenAI or return mock text
	return fmt.Sprintf("[OpenAI] Response for prompt: %s", prompt), nil
}

// -------------------
// Local LLM Implementation
// -------------------
type LocalLLMClient struct {}

func NewLocalLLMClient() *LocalLLMClient {
	return &LocalLLMClient{}
}

func (c *LocalLLMClient) Generate(prompt string) (string, error) {
	// Placeholder: call local LLM binary or model
	return fmt.Sprintf("[Local LLM] Response for prompt: %s", prompt), nil
}


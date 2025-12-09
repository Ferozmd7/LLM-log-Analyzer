package controllers

import (
	"os/exec"
	"strings"
)

// OllamaClient implements LLMClient using the Ollama CLI
type OllamaClient struct {
	Model string
}

// NewOllamaClient creates a new Ollama client
func NewOllamaClient(model string) *OllamaClient {
	return &OllamaClient{Model: model}
}

// Generate runs Ollama CLI to get a response
func (c *OllamaClient) Generate(prompt string) (string, error) {
	cmd := exec.Command("ollama", "run", c.Model, prompt)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}


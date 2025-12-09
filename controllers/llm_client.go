package controllers

// LLMClient defines a generic interface for LLM backends
type LLMClient interface {
	Generate(prompt string) (string, error)
}


package controllers

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    openai "github.com/openai/openai-go"
)

type LLMAnalysis struct {
    Summary         string   `json:"summary"`
    Anomalies       []string `json:"anomalies"`
    Recommendations []string `json:"recommendations"`
}

type LLMClient struct {
    OpenAI *openai.Client
}

func NewLLMClient() *LLMClient {
    return &LLMClient{
        OpenAI: openai.NewClient(os.Getenv("OPENAI_API_KEY")),
    }
}

func (c *LLMClient) AnalyzeLogs(ctx context.Context, model string, logs string) (*LLMAnalysis, error) {
    resp, err := c.OpenAI.Chat.Completions.Create(ctx, openai.ChatCompletionRequest{
        Model: model,
        Messages: []openai.ChatCompletionMessage{
            {Role: "system", Content: "Return JSON: {summary, anomalies[], recommendations[]}"},
            {Role: "user", Content: logs},
        },
    })
    if err != nil {
        return nil, err
    }

    var out LLMAnalysis
    if err := json.Unmarshal([]byte(resp.Choices[0].Message.Content), &out); err != nil {
        return nil, fmt.Errorf("invalid JSON: %w", err)
    }

    return &out, nil
}

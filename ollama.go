package main

// ChatRequest represents a chat request from the client.
type ChatRequest struct {
	Model  string `json:"model" form:"model"`
	Prompt string `json:"prompt" form:"prompt"`
}

// ChatResponse represents the final Ollama API response payload.
type ChatResponse struct {
	Model              string `json:"model"`
	CreatedAt          string `json:"created_at"`
	Response           string `json:"response"`
	Done               bool   `json:"done"`
	Context            []int  `json:"context"`
	TotalDuration      int64  `json:"total_duration"`
	LoadDuration       int64  `json:"load_duration"`
	PromptEvalCount    int    `json:"prompt_eval_count"`
	PromptEvalDuration int64  `json:"prompt_eval_duration"`
	EvalCount          int    `json:"eval_count"`
	EvalDuration       int64  `json:"eval_duration"`
}

// OllamaStreamChunk matches the streaming JSON lines returned by the Ollama API.
type OllamaStreamChunk struct {
	Type     string `json:"type,omitempty"`
	Response string `json:"response,omitempty"`
	Text     string `json:"text,omitempty"`
	Done     bool   `json:"done,omitempty"`
}

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type OllamaService struct {
	endpoint string
	model    string
	client   *http.Client
}

// NewOllamaService returns a new service configured for the local Ollama endpoint.
func NewOllamaService(endpoint, model string) *OllamaService {
	return &OllamaService{
		endpoint: endpoint,
		model:    model,
		client:   &http.Client{},
	}
}

// StreamChat sends a streaming request to Ollama and returns the raw HTTP response.
func (s *OllamaService) StreamChat(prompt string) (*http.Response, error) {
	requestBody := map[string]interface{}{
		"model":  s.model,
		"prompt": prompt,
		"stream": true,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, s.endpoint, bytes.NewReader(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to build request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to call Ollama: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		return nil, fmt.Errorf("ollama returned %d: %s", resp.StatusCode, string(body))
	}

	return resp, nil
}

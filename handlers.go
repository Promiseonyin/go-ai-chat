package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Handler struct {
	service *OllamaService
}

// NewHandler creates a new handler with the Ollama service dependency.
func NewHandler(service *OllamaService) *Handler {
	return &Handler{service: service}
}

// HandleIndex serves the chat UI.
func (h *Handler) HandleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	http.ServeFile(w, r, "index.html")
}

// HandleChat receives the prompt and streams AI output as SSE.
func (h *Handler) HandleChat(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var prompt string
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Unable to parse form", http.StatusBadRequest)
			return
		}
		prompt = r.FormValue("prompt")
	} else {
		prompt = r.URL.Query().Get("prompt")
	}

	prompt = strings.TrimSpace(prompt)
	if prompt == "" {
		http.Error(w, "Prompt cannot be empty", http.StatusBadRequest)
		return
	}

	response, err := h.service.StreamChat(prompt)
	if err != nil {
		http.Error(w, fmt.Sprintf("Ollama service error: %v", err), http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.WriteHeader(http.StatusOK)

	reader := bufio.NewReader(response.Body)
	prevText := ""
	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			event := fmt.Sprintf("data: %v\n\n", err)
			fmt.Fprint(w, event)
			flusher.Flush()
			return
		}

		line = strings.TrimSpace(line)
		if line == "" {
			if err == io.EOF {
				break
			}
			continue
		}

		var chunk OllamaStreamChunk
		if err := json.Unmarshal([]byte(line), &chunk); err != nil {
			continue
		}

		content := chunk.Response
		if content == "" {
			content = chunk.Text
		}

		delta := content
		if strings.HasPrefix(content, prevText) {
			delta = content[len(prevText):]
		}

		if delta != "" {
			streamTextAsSSE(w, flusher, delta)
		}
		prevText = content

		if chunk.Done {
			sendDone(w, flusher)
			return
		}

		if err == io.EOF {
			break
		}
	}

	sendDone(w, flusher)
}

func streamTextAsSSE(w http.ResponseWriter, flusher http.Flusher, text string) {
	fmt.Fprintf(w, "data: %s\n\n", text)
	flusher.Flush()
}

func sendDone(w http.ResponseWriter, flusher http.Flusher) {
	fmt.Fprint(w, "event: done\ndata: true\n\n")
	flusher.Flush()
}

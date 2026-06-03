package main

import (
	"log"
	"net/http"
)

func main() {
	service := NewOllamaService("http://localhost:11434/api/generate", "llama3.2")
	handler := NewHandler(service)

	http.HandleFunc("/", handler.HandleIndex)
	http.HandleFunc("/chat", handler.HandleChat)

	log.Println("Starting AI chat server on http://localhost:8082")
	if err := http.ListenAndServe(":8082", nil); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}

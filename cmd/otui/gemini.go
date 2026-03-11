package main

import (
  "context"
  "fmt"
  "log"
  "google.golang.org/genai"
)

func GeminiClient(ctx context.Context) (*genai.Client) {
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

// func callGemini() {
func main(){
  ctx := context.Background()

	client := GeminiClient(ctx)

  history := []*genai.Content{}

  chat, _ := client.Chats.Create(ctx, "gemini-3.1-flash-lite-preview", nil, history)
  stream := chat.SendMessageStream(ctx, genai.Part{Text: "How do I use the bubbletea framework?"})

  for chunk, _ := range stream {
      part := chunk.Candidates[0].Content.Parts[0]
      fmt.Print(part.Text)
  }
}
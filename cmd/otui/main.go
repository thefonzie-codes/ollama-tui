package main

import (
	"fmt"
	"log"
	"os"

	tea "charm.land/bubbletea/v2"
	dotenv "github.com/joho/godotenv"
)

var url string = "http://localhost:11434"

func main() {
	// Enable Error Logging
	logError()
	err := dotenv.Load()

	if err != nil {
		fmt.Printf("Error loading dotenv: %v", err)
	}

	url = os.Getenv("OLLAMA_URL")

	// uncomment below if default ollama URL is different

	// if url == "http://localhost:11434" {
	//
	// }

	fmt.Println(url)

	if !IsOllamaRunning() {
		fmt.Printf("IsOllamaRunning = %v \n", IsOllamaRunning())
		fmt.Println("Oops, looks like Ollama is not running...")
		StartOllama()
	}

	// prints "Hello" for now
	p := tea.NewProgram(initialModel())

	if _, err := p.Run(); err != nil {
		log.Fatalf("Womp womp, error starting Bubbletea: %v \n", err)
		os.Exit(1)
	}

	models, err := GetOllamaModels()
	if err != nil {
		fmt.Printf("Womp womp, couldn't get the Ollama models: %v \n", err)
		os.Exit(1)
	}

	fmt.Println(models)
}

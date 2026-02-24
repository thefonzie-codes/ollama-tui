package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	dotenv "github.com/joho/godotenv"
)

func initialModel() model {
	return model("Hello")
}

// var URL string = os.Getenv("OLLAMA_URL")
var url string = "http://localhost:11434"

func main() {

	err := dotenv.Load()

	if err != nil {
		fmt.Printf("Error loading dotenv: %v", err)
	}

	fmt.Println(url)

	if !IsOllamaRunning() {
		fmt.Printf("IsOllamaRunning = %v \n", IsOllamaRunning())
		fmt.Println("Oops, looks like Ollama is not running...")
		StartOllama()
	}

	// prints "Hello" for now
	p := tea.NewProgram(initialModel())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Womp womp, there's an error: %v \n", err)
		os.Exit(1)
	}

	models, err := CallOllama()
	if err != nil {
		fmt.Printf("Womp womp, there's an error: %v \n", err)
		os.Exit(1)
	}

	fmt.Println(models)
}

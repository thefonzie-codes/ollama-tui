package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"time"
)

type OllamaListResponse struct {
	Models []OllamaModelInfo `json:"models"`
}

type OllamaModelInfo struct {
	Name string `json:"name"`
	// Size    string             `json:"size"`
	// Details OllamaModelDetails `json:"details"`
}

// type OllamaModelDetails struct {
// 	Family string `json:"family"`
// }

// Quick check to see if Ollama is running by calling the /api/tags endpoint

func IsOllamaRunning() bool {

	client := http.Client{
		Timeout: 2 * time.Second,
	}

	fmt.Printf("URL: %v \n", url)

	resp, err := client.Get(url + "/api/tags")

	if err != nil {
		return false
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("error: %s", resp.Status)
	}

	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK
}

// Start Ollama if it is not running

func StartOllama() {

	fmt.Println("Attempting to start Ollama...")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cmd := exec.Command("ollama", "serve")

	go func() {
		if err := cmd.Run(); err != nil && ctx.Err() == nil {
			fmt.Println("Failed to start Ollama in 'StartOllama()'", err)
			cancel()
		}

	}()

	time.Sleep(2 * time.Second)

	if !IsOllamaRunning() {
		fmt.Println("Ollama failed to start")
		os.Exit(1)
	}

	fmt.Println("Ollama started successfully")
}

// Call Ollama to get a list of models

func GetOllamaModels() (OllamaListResponse, error) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", url+"/api/tags", nil)

	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return OllamaListResponse{}, fmt.Errorf("error: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return OllamaListResponse{}, err
	}

	var listResponse OllamaListResponse

	if err := json.Unmarshal(body, &listResponse); err != nil {
		return OllamaListResponse{}, err
	}

	fmt.Printf("Models: %v \n", listResponse)
	return listResponse, nil
}

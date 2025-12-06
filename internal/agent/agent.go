package main

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/adk/agent/llmagent"
	"google.golang.org/adk/cmd/launcher/adk"
	"google.golang.org/adk/cmd/launcher/full"
	"google.golang.org/adk/model/gemini"
	"google.golang.org/adk/server/restapi/services"
	"google.golang.org/adk/tool"
	"google.golang.org/adk/tool/geminitool"
	"google.golang.org/genai"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	ctx := context.Background()

	model, err := gemini.NewModel(ctx, "gemini-2.5-flash", &genai.ClientConfig{
		APIKey: os.Getenv("GOOGLE_API_KEY"),
	})
	if err != nil {
		fmt.Printf("Failed to create model: %v", err)
		return
	}

	agent, err := llmagent.New(llmagent.Config{
		Name:        "PIM Agent",
		Model:       model,
		Description: "An agent that helps manage personal information and tasks.",
		Tools: []tool.Tool{
			geminitool.GoogleSearch{},
		},
	})
	if err != nil {
		fmt.Printf("Failed to create agent: %v", err)
		return
	}

	config := &adk.Config{
		AgentLoader: services.NewSingleAgentLoader(agent),
	}
	l := full.NewLauncher()
	if err := l.Execute(ctx, config, os.Args[1:]); err != nil {
		fmt.Printf("Failed to run launcher: %v", err)
	}
}

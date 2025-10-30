package main

import (
	"context"
	"fmt"
	"time"

	"github.com/joho/godotenv"

	"github.com/nlpodyssey/openai-agents-go/agents"
	"github.com/nlpodyssey/openai-agents-go/modelsettings"
)

func main() {
	_ = godotenv.Load()

	// We'll run the Streamable HTTP server in a goroutine. Usually this would be a remote server,
	// but for this demo, we'll run it locally at http://localhost:8000
	// Proper server handling, including graceful shutdown, is omitted just to keep the example short.
	go runServer("localhost:8000")

	time.Sleep(2 * time.Second) // Give it 2 seconds to start

	server := agents.NewMCPServerStreamableHTTP(agents.MCPServerStreamableHTTPParams{
		Name: "Streamable HTTP Go Server",
		URL:  "http://localhost:8000",
	})

	err := server.Run(context.Background(), func(ctx context.Context, server *agents.MCPServerWithClientSession) error {
		return run(ctx, server)
	})
	if err != nil {
		panic(err)
	}
}

func run(ctx context.Context, mcpServer agents.MCPServer) error {
	agent := agents.New("Assistant").
		WithInstructions("Use the tools to answer the questions.").
		AddMCPServer(mcpServer).
		WithModelSettings(modelsettings.ModelSettings{
			ToolChoice: modelsettings.ToolChoiceRequired,
		}).
		WithModel("gpt-4o")

	// Use the `add` tool to add two numbers
	message := "Add these numbers: 7 and 22."
	fmt.Println("Running:", message)
	result, err := agents.Run(ctx, agent, message)
	if err != nil {
		return err
	}
	fmt.Println(result.FinalOutput)

	// Run the `get_weather` tool
	message = "What's the weather in Tokyo?"
	fmt.Println("\nRunning:", message)
	result, err = agents.Run(ctx, agent, message)
	if err != nil {
		return err
	}
	fmt.Println(result.FinalOutput)

	// Run the `get_secret_word` tool
	message = "What's the secret word?"
	fmt.Println("\nRunning:", message)
	result, err = agents.Run(ctx, agent, message)
	if err != nil {
		return err
	}
	fmt.Println(result.FinalOutput)

	return nil
}

package main

import (
	"context"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func runServer(addr string) {
	server := mcp.NewServer(&mcp.Implementation{Name: "Agents MCP Streamable HTTP Example"}, nil)

	type addParams struct {
		A int `json:"a"`
		B int `json:"b"`
	}

	type addResult struct {
		Result int `json:"result"`
	}

	mcp.AddTool(
		server, &mcp.Tool{Name: "add", Description: "Add two numbers"},
		func(_ context.Context, _ *mcp.CallToolRequest, params addParams) (*mcp.CallToolResult, addResult, error) {
			fmt.Printf("[debug-server] add(%d, %d)\n", params.A, params.B)
			result := params.A + params.B
			return &mcp.CallToolResult{
				Content: []mcp.Content{
					&mcp.TextContent{Text: strconv.Itoa(result)},
				},
			}, addResult{Result: result}, nil
		},
	)

	type getSecretWordResult struct {
		SecretWord string
	}

	mcp.AddTool(
		server, &mcp.Tool{Name: "get_secret_word"},
		func(_ context.Context, _ *mcp.CallToolRequest, _ struct{}) (*mcp.CallToolResult, getSecretWordResult, error) {
			fmt.Println("[debug-server] get_secret_word()")
			choice := []string{"apple", "banana", "cherry"}[rand.Intn(3)]
			return &mcp.CallToolResult{
				Content: []mcp.Content{
					&mcp.TextContent{Text: choice},
				},
			}, getSecretWordResult{SecretWord: choice}, nil
		},
	)

	type getCurrentWeatherParams struct {
		City string `json:"city"`
	}

	type getCurrentWeatherResult struct {
		Weather string
	}

	mcp.AddTool(
		server, &mcp.Tool{Name: "get_current_weather"},
		func(ctx context.Context, _ *mcp.CallToolRequest, params getCurrentWeatherParams) (*mcp.CallToolResult, *getCurrentWeatherResult, error) {
			fmt.Printf("[debug-server] get_current_weather(%q)\n", params.City)

			resp, err := http.Get("https://wttr.in/" + params.City)
			if err != nil {
				return nil, nil, fmt.Errorf("HTTP request to wttr.in error: %w", err)
			}
			defer func() {
				_ = resp.Body.Close()
			}()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, nil, fmt.Errorf("failed to read wttr.in body response: %w", err)
			}

			return &mcp.CallToolResult{
				Content: []mcp.Content{
					&mcp.TextContent{Text: string(body)},
				},
			}, &getCurrentWeatherResult{Weather: string(body)}, nil
		},
	)

	handler := mcp.NewStreamableHTTPHandler(func(*http.Request) *mcp.Server {
		return server
	}, nil)

	fmt.Printf("Starting Streamable HTTP server at %s ...\n", addr)
	err := http.ListenAndServe(addr, handler)
	if err != nil {
		panic(err)
	}
}

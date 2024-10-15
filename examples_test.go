package goopenrouter_test

import (
	"context"
	"errors"
	"fmt"

	"github.com/affirm-bats-yodel/goopenrouter"
)

func ExampleNewClient() {
	var openRouterAPIKey = "api-key"
	client := goopenrouter.NewClient(openRouterAPIKey)

	_ = client
}

func ExampleNewClientFromEnv() {
	// First, you should've Export a
	// API Key to Environment Variable
	//
	// Using `export` (or use `$env:` on Windows PowerShell)
	//
	// on Linux:
	//   $ export OPENROUTER_API_KEY="${API_KEY}"
	//
	// on Windows:
	//   $ $env:OPENROUTER_API_KEY="${API_KEY}"

	// If There's No Value were found on `OPENROUTER_API_KEY`
	// It'll return an error as ErrEnvNoRouterKey

	client, err := goopenrouter.NewClientFromEnv()
	if err != nil && errors.Is(err, goopenrouter.ErrEnvNoRouterKey) {
		panic("todo: handle error")
	} else if err != nil {
		panic("panic")
	}

	_ = client
}

func ExampleClient_GetLimits() {
	client, err := goopenrouter.NewClientFromEnv()
	if err != nil {
		panic(err) // handle error gracefully...
	}

	// You can Retrieve a Status of API Account
	limits, err := client.GetLimits(context.Background())
	if err != nil {
		panic(err) // handle error gracefully...
	}

	fmt.Println(limits)
}

func ExampleClient_GetModels() {
	// Create a Client from Env or Plain
	client, err := goopenrouter.NewClientFromEnv()
	if err != nil {
		panic(err)
	}

	// you can filter a supported model
	// using parameters option.
	//
	// for example, let's use top_p
	supportedParameters := []string{
		"top_p",
	}

	// then let's call GetModels to list models
	models, err := client.GetModels(
		context.Background(),
		supportedParameters...,
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(models)
}

func ExampleClient_GetParameters() {
	client, err := goopenrouter.NewClientFromEnv()
	if err != nil {
		panic(err)
	}

	// Let's Retrieve a Parameters of
	// "openai/chatgpt-4o-latest"
	//
	// https://openrouter.ai/openai/chatgpt-4o-latest

	// You can Copy a Model ID from Web Dashboard
	modelID := "openai/chatgpt-4o-latest"

	// Provider is OpenAI
	provider := "OpenAI"

	params, err := client.GetParameters(
		context.Background(),
		modelID,
		provider,
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(params)
}

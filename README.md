# goopenrouter 

[![Go Reference](https://pkg.go.dev/badge/github.com/affirm-bats-yodel/goopenrouter.svg)](https://pkg.go.dev/github.com/affirm-bats-yodel/goopenrouter)

A [OpenRouter](https://openrouter.ai) Golang Client

## Why?

I've started to use OpenRouter, however there's no ways 
to list models, limits and parameters via Go.

## Supported Endpoints

### OpenRouter Models API

> https://openrouter.ai/docs/models

* GoDoc Example: [GetModel](https://pkg.go.dev/github.com/affirm-bats-yodel/goopenrouter#example-Client.GetModels)

### OpenRouter Limits API

> https://openrouter.ai/docs/limits

* GoDoc Example: [GetLimits](https://pkg.go.dev/github.com/affirm-bats-yodel/goopenrouter#example-Client.GetLimits)

### OpenRouter Parameters API

> https://openrouter.ai/docs/parameters-api

* GoDoc Example: [GetParameters](https://pkg.go.dev/github.com/affirm-bats-yodel/goopenrouter#example-Client.GetParameters)

## Unsupported Methods or Limitation

### Completions API

* Use OpenAI compatible SDKs.

## Having Trouble?

* You can Raise a New Issue or make a Pull Request.
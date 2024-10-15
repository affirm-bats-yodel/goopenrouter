package openrouter

import "context"

// ClientInterface Interface for OpenRouter Client
//
// - Define a functions to implement
type ClientInterface interface {
	Models(ctx context.Context, parameters ...string)
}

package openrouter

import "context"

// ClientInterface Interface for OpenRouter Client
//
// - Define a functions to implement
type ClientInterface interface {
	// GetModels Get Model Information
	GetModels(ctx context.Context, parameters ...string) ([]*Model, error)
	// GetLimits Get Limit Information
	GetLimits(ctx context.Context) (*Limit, error)
	// GetParameters Get Parameter Information
	//
	// Sometimes, modelID and provider can contain a slash ("/")
	// or Whitespaces (" "). Handler MUST BE Handled.
	GetParameters(ctx context.Context, modelID string, provider string) (*Parameters, error)
}

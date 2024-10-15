package openrouter

import "context"

// ClientInterface Interface for OpenRouter Client
//
// - Define a functions to implement
type ClientInterface interface {
	// GetModels Get Model Information
	GetModels(ctx context.Context, parameters ...string)
	// GetLimits Get Limit Information
	GetLimits(ctx context.Context) (*Limit, error)
}

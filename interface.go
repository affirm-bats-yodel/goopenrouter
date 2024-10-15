package goopenrouter

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
	//
	// - modelID: Specify a model ID, since optimal parameters vary between models.
	//
	// - providers: Specify a provider name to filter supported parameters
	// for a given model by provider.
	GetParameters(ctx context.Context, modelID string, providers ...string) (*Parameters, error)
}

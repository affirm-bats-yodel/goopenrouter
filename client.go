package openrouter

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"os"
)

// EnvOpenRouterKey a Environment Key for OpenRouter API Key
const EnvOpenRouterKey = "OPENROUTER_API_KEY"

var ErrEnvNoRouterKey = errors.New("error: no API key exist on $" + EnvOpenRouterKey)

// NewClient Create new OpenRouter Client
func NewClient() (*Client, error) {
	v := os.Getenv(EnvOpenRouterKey)
	if v == "" {
		return nil, ErrEnvNoRouterKey
	}
	return &Client{
		APIKey: v,
	}, nil
}

// Client OpenRouter Client Implementation
type Client struct {
	APIKey string
}

// GetLimits implements ClientInterface.
func (c *Client) GetLimits(ctx context.Context) (*Limit, error) {
	var (
		lr Response[*Limit]
	)

	req, err := newHTTPRequest(ctx, "GET", "/api/v1/auth/key", c.APIKey)
	if err != nil {
		return nil, err
	}

	res, err := newHTTPClient().Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&lr); err != nil {
		return nil, err
	}

	if err := res.Body.Close(); err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, lr.Error
	}

	return lr.Data, nil
}

// GetModels implements ClientInterface.
func (c *Client) GetModels(ctx context.Context, parameters ...string) ([]*Model, error) {
	var lr Response[[]*Model]

	req, err := newHTTPRequest(ctx, "GET", "/api/v1/models", c.APIKey)
	if err != nil {
		return nil, err
	}

	res, err := newHTTPClient().Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&lr); err != nil {
		return nil, err
	}

	if err := res.Body.Close(); err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, lr.Error
	}

	return lr.Data, nil
}

var _ ClientInterface = (*Client)(nil)

// newHTTPClient Create a new http.Client
func newHTTPClient() *http.Client {
	return &http.Client{}
}

const (
	openRouterAddr      = "https://openrouter.ai"
	authorizationHeader = "Authorization"
)

// newHTTPRequest Create a http.Request
func newHTTPRequest(ctx context.Context, method string, endpoint string, apiKey string) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, openRouterAddr+endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(authorizationHeader, "Bearer "+apiKey)
	return req, nil
}

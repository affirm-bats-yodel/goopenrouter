package openrouter

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"os"
)

// EnvOpenRouterKey a Environment Key for OpenRouter API Key
const EnvOpenRouterKey = "OPENROUTER_API_KEY"

var ErrEnvNoRouterKey = errors.New("error: no API key exist on $" + EnvOpenRouterKey)

// NewClient Create a New Client
func NewClient(apiKey string) *Client {
	return &Client{
		APIKey: apiKey,
	}
}

// NewClientFromEnv Create a New Client
//
// Check Environment Variable $OPENROUTER_API_KEY
// and it'll return an error if it does not exist.
//
// If Exist, value will be used for Authentication
func NewClientFromEnv() (*Client, error) {
	v := os.Getenv(EnvOpenRouterKey)
	if v == "" {
		return nil, ErrEnvNoRouterKey
	}
	return NewClient(v), nil
}

// Client OpenRouter Client Implementation
//
// TODO: Refactor to minimize Duplicated Codes
type Client struct {
	APIKey string
}

// GetLimits implements ClientInterface.
func (c *Client) GetLimits(ctx context.Context) (*Limit, error) {
	return doRequest[*Limit](ctx, "GET", "/api/v1/auth/key", c.APIKey)
}

// GetModels implements ClientInterface.
func (c *Client) GetModels(ctx context.Context, parameters ...string) ([]*Model, error) {
	return doRequest[[]*Model](ctx, "GET", "/api/v1/models", c.APIKey)
}

// GetParameters implements ClientInterface.
func (c *Client) GetParameters(ctx context.Context, modelID string, providers ...string) (*Parameters, error) {
	var (
		endpoint = "/api/v1/parameters"
		uvs      url.Values
	)

	endpoint = endpoint + url.QueryEscape(modelID)

	if providers != nil && providers[0] != "" {
		uvs.Add("provider", url.QueryEscape(providers[0]))
	}

	return doRequest[*Parameters](ctx, "GET", endpoint, c.APIKey, uvs)
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

func doRequest[data any](ctx context.Context, method, endpoint, apiKey string, queryParams ...url.Values) (data, error) {
	var body Response[data]

	req, err := http.NewRequestWithContext(ctx, method, openRouterAddr+endpoint, nil)
	if err != nil {
		return body.Data, nil
	}
	req.Header.Add(authorizationHeader, "Bearer "+apiKey)

	if len(queryParams) > 0 {
		req.URL.RawQuery = queryParams[0].Encode()
	}

	res, err := newHTTPClient().Do(req)
	if err != nil {
		return body.Data, err
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		return body.Data, err
	}

	if err := res.Body.Close(); err != nil {
		return body.Data, err
	}

	if res.StatusCode != http.StatusOK {
		return body.Data, body.Error
	}

	return body.Data, nil
}

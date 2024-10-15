package goopenrouter

import "time"

type Response[data any] struct {
	Data  data         `json:"data,omitempty"`
	Error *ErrorDetail `json:"error,omitempty"`
}

// Limit Information about Request Limit
type Limit struct {
	Label string `json:"label"`
	// Usage Number of credits used
	Usage float64 `json:"usage"`
	// Limit Credit limit for the key
	// or null if unlimited
	Limit *int64 `json:"limit"`
	// IsFreeTier Whether the user has paid
	// for credits before
	IsFreeTier bool `json:"is_free_tier"`
	// RateLimit Rate Limit information
	RateLimit LimitRateLimit `json:"rate_limit"`
}

// LimitRateLimit Information about Rate Limit
type LimitRateLimit struct {
	// Requests Number of requests allowed
	Requests int `json:"requests"`
	// Interval in this interval, e.g. "10s"
	Interval string `json:"interval"`
}

// GetInterval Get Interval as Go's time.Duration
func (r *LimitRateLimit) GetInterval() time.Duration {
	v, err := time.ParseDuration(r.Interval)
	if err != nil {
		return 0
	}
	return v
}

// Model Information about LLM Model
type Model struct {
	ID               string                 `json:"id"`
	Name             string                 `json:"name"`
	Created          int64                  `json:"created"`
	Description      string                 `json:"description"`
	Pricing          *ModelPricing          `json:"pricing"`
	ContextLength    int64                  `json:"context_length"`
	Architecture     *ModelArchitecture     `json:"architecture"`
	TopProvider      *ModelTopProvider      `json:"top_provider"`
	PerRequestLimits *ModelPerRequestLimits `json:"per_request_limits"`
}

type ModelPricing struct {
	Prompt     string `json:"prompt"`
	Completion string `json:"completion"`
	Request    string `json:"request"`
	Image      string `json:"image"`
}

type ModelArchitecture struct {
	Tokenizer    string `json:"tokenizer"`
	InstructType string `json:"instruct_type"`
	Modality     string `json:"modality"`
}

type ModelTopProvider struct {
	ContextLength       int64 `json:"context_length"`
	MaxCompletionTokens int64 `json:"max_completion_tokens"`
	IsModerated         bool  `json:"is_moderated"`
}

type ModelPerRequestLimits struct {
	PromptTokens     string `json:"prompt_tokens"`
	CompletionTokens string `json:"completion_tokens"`
}

type Parameters struct {
	Model                string   `json:"model"`
	SupportedParameters  []string `json:"supported_parameters"`
	FrequencyPenaltyP10  float64  `json:"frequency_penalty_p10"`
	FrequencyPenaltyP50  float64  `json:"frequency_penalty_p50"`
	FrequencyPenaltyP90  float64  `json:"frequency_penalty_p90"`
	MinPP10              float64  `json:"min_p_p10"`
	MinPP50              float64  `json:"min_p_p50"`
	MinPP90              float64  `json:"min_p_p90"`
	PresencePenaltyP10   float64  `json:"presence_penalty_p10"`
	PresencePenaltyP50   float64  `json:"presence_penalty_p50"`
	PresencePenaltyP90   float64  `json:"presence_penalty_p90"`
	RepetitionPenaltyP10 float64  `json:"repetition_penalty_p10"`
	RepetitionPenaltyP50 float64  `json:"repetition_penalty_p50"`
	RepetitionPenaltyP90 float64  `json:"repetition_penalty_p90"`
	TemperatureP10       float64  `json:"temperature_p10"`
	TemperatureP50       float64  `json:"temperature_p50"`
	TemperatureP90       float64  `json:"temperature_p90"`
	TopAP10              float64  `json:"top_a_p10"`
	TopAP50              float64  `json:"top_a_p50"`
	TopAP90              float64  `json:"top_a_p90"`
	TopKP10              float64  `json:"top_k_p10"`
	TopKP50              float64  `json:"top_k_p50"`
	TopKP90              float64  `json:"top_k_p90"`
	TopPP10              float64  `json:"top_p_p10"`
	TopPP50              float64  `json:"top_p_p50"`
	TopPP90              float64  `json:"top_p_p90"`
}

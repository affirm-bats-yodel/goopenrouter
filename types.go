package openrouter

import "time"

// LimitResponse Limit Response
type LimitResponse struct {
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
	RateLimit RateLimit `json:"rate_limit"`
}

// RateLimit Information about Rate Limit
type RateLimit struct {
	// Requests Number of requests allowed
	Requests int `json:"requests"`
	// Interval in this interval, e.g. "10s"
	Interval string `json:"interval"`
}

// GetInterval Get Interval as Go's time.Duration
func (r *RateLimit) GetInterval() time.Duration {
	v, err := time.ParseDuration(r.Interval)
	if err != nil {
		return 0
	}
	return v
}

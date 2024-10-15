package goopenrouter

import "fmt"

// StatusCode OpenRouter Status Code
//
// refer from OpenRouter's Errors Page
//
// https://openrouter.ai/docs/errors#error-codes
type StatusCode int

const (
	// StatusBadRequest Bad Request
	//
	// Invalid or missing params, CORS
	StatusBadRequest StatusCode = 400
	// StatusInvalidCredentials Invalid credentials
	//
	// OAuth session expired, disabled/invalid API key
	StatusInvalidCredentials StatusCode = 401
	// StatusInsufficientCredit Insufficient Credit
	//
	// Your account or API key has insufficient credits.
	// Add more credits and retry the request.
	StatusInsufficientCredit StatusCode = 402
	// StatusFlagged Input Flagged
	//
	// Your chosen model requires moderation
	// and your input was flagged
	StatusFlagged StatusCode = 403
	// StatusTimedout Timed Out
	//
	// Your request timed out
	StatusTimedout StatusCode = 408
	// StatusRateLimited Rate Limited
	//
	// You are being rate limited
	StatusRateLimited StatusCode = 429
	// StatusInvalidResponse Invalid Response from Provider
	//
	// Your chosen model is down
	// or we received an invalid response from it
	StatusInvalidResponse StatusCode = 502
	// StatusUnavailable Model Unavailable
	//
	// There is no available model provider that
	// meets your routing requirements
	StatusUnavailable StatusCode = 503
)

// ErrorDetail Error Detail
type ErrorDetail struct {
	// Code Error Code
	Code StatusCode `json:"code"`
	// Message Error Message
	Message string `json:"message"`
}

// Error implements error.
func (e *ErrorDetail) Error() string {
	return fmt.Sprintf("error (%d): %s", e.Code, e.Message)
}

var _ error = (*ErrorDetail)(nil)

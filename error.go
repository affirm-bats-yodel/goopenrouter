package openrouter

import "fmt"

// IsErrorResponse Check error is ErrorResponse
func IsErrorResponse(e error) bool {
	if e == nil {
		return false
	}
	_, b := e.(*ErrorResponse)
	return b
}

// ErrorResponse Error Response from OpenRouter
//
// Compatible with Golang's Error interface.
//
// For Example:
//
//	{"error": {"code": 0, "message": "string"}}
type ErrorResponse struct {
	ErrorDetail *ErrorDetail `json:"error"`
}

// UnknownError Error Message when there is no
// message (or if ErrorDetail is Empty)
const UnknownError = "error: unknown error occurred"

// Error implements error.
func (e *ErrorResponse) Error() string {
	if e.ErrorDetail != nil && e.ErrorDetail.Message != "" {
		return fmt.Sprintf("error: (%d): %s", e.ErrorDetail.Code, e.ErrorDetail.Message)
	}
	return UnknownError
}

var _ error = (*ErrorResponse)(nil)

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

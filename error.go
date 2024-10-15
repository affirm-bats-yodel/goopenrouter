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

// ErrorDetail Error Detail
type ErrorDetail struct {
	// Code Error Code
	Code int `json:"code"`
	// Message Error Message
	Message string `json:"message"`
}

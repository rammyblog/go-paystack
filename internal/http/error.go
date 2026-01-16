package http

import (
	"encoding/json"
	"net/http"
	"net/url"
)

// APIError includes the response from the Paystack API and some HTTP request info
type APIError struct {
	Message        string        `json:"message,omitempty"`
	HTTPStatusCode int           `json:"code,omitempty"`
	Details        ErrorResponse `json:"details,omitempty"`
	URL            *url.URL      `json:"url,omitempty"`
	Header         http.Header   `json:"header,omitempty"`
}

// APIError supports the error interface
func (e *APIError) Error() string {
	ret, _ := json.Marshal(e)
	return string(ret)
}

// ErrorResponse represents an error response from the Paystack API server
type ErrorResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message,omitempty"`
}

func NewAPIError(resp *http.Response, data map[string]interface{}) *APIError {
	var status bool
	var message string

	if s, ok := data["status"].(bool); ok {
		status = s
	}

	if m, ok := data["message"].(string); ok {
		message = m
	}

	response := ErrorResponse{
		Status:  status,
		Message: message,
	}
	return &APIError{
		HTTPStatusCode: resp.StatusCode,
		Header:         resp.Header,
		Details:        response,
		URL:            resp.Request.URL,
	}
}

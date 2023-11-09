// From https://github.dev/rpip/paystack-go/blob/master/paystack.go

package paystack

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

func newAPIError(resp *http.Response, data map[string]interface{}) *APIError {
	response := ErrorResponse{
		Status:  data["status"].(bool),
		Message: data["message"].(string),
	}
	return &APIError{
		HTTPStatusCode: resp.StatusCode,
		Header:         resp.Header,
		Details:        response,
		URL:            resp.Request.URL,
	}
}

package paystack

import (
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"time"
)

type Client struct {
	APIKey      string
	HttpClient  *http.Client
	log         *Logger
	Transaction *Transaction
	BaseUrl     *url.URL
}

const BASE_URL = "https://api.paystack.co"

// Response represents arbitrary response data
type APIResponse map[string]interface{}

type Metadata map[string]interface{}

// PaginationMeta is pagination metadata for paginated responses from the Paystack API
type PaginationMeta struct {
	Total     int `json:"total"`
	Skipped   int `json:"skipped"`
	PerPage   int `json:"perPage"`
	Page      int `json:"page"`
	PageCount int `json:"pageCount"`
}

// NewClient creates a new Paystack API client with the given API key.
func NewClient(apiKey string) *Client {
	httpClient := &http.Client{
		Timeout: 5 * time.Second,
	}
	log := NewLogger(slog.New(slog.NewTextHandler(os.Stdout, nil)))
	parsedUrl, _ := url.Parse(BASE_URL)
	c := &Client{APIKey: apiKey, HttpClient: httpClient, log: log, BaseUrl: parsedUrl}
	c.Transaction = newTransaction(c)

	return c
}

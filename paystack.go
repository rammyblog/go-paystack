package paystack

import (
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type Client struct {
	APIKey           string
	HttpClient       *http.Client
	log              *slog.Logger
	Transaction      *Transaction
	TransactionSplit *TransactionSplit
	Plan             *Plans
	BaseUrl          *url.URL
}

const BASE_URL = "https://api.paystack.co"

// Response represents arbitrary response data
type APIResponse map[string]interface{}

type GenericResponse map[string]interface{}

type Metadata map[string]interface{}

type QueryType struct {
	Key   string
	Value string
}

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

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	parsedUrl, _ := url.Parse(BASE_URL)
	c := &Client{APIKey: apiKey, HttpClient: httpClient, log: logger, BaseUrl: parsedUrl}
	c.Transaction = newTransaction(c)
	c.TransactionSplit = newTransactionSplit(c)
	c.Plan = newPlans(c)

	return c
}

func Query(key, value string) QueryType {
	return QueryType{
		Key:   key,
		Value: value,
	}
}

func addQueryToUrl(url string, queries ...QueryType) string {
	for _, query := range queries {
		if strings.Contains(url, "?") {
			url += fmt.Sprintf("&%s=%s", query.Key, query.Value)
		} else {
			url += fmt.Sprintf("?%s=%s", query.Key, query.Value)
		}
	}
	return url
}

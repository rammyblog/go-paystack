package paystack

import (
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/rammyblog/go-paystack/internal/types"
	"github.com/rammyblog/go-paystack/plans"
	"github.com/rammyblog/go-paystack/subscription"
	"github.com/rammyblog/go-paystack/transaction"
	transaction_splits "github.com/rammyblog/go-paystack/transaction-splits"
)

const BASE_URL = "https://api.paystack.co"

type Client struct {
	APIKey           string
	HttpClient       *http.Client
	Log              *slog.Logger
	Transaction      *transaction.Transaction
	TransactionSplit *transaction_splits.TransactionSplit
	Plan             *plans.Plans
	Subscription     *subscription.Subscription
	BaseUrl          *url.URL
}

type ClientOption func(*Client)

func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) {
		c.HttpClient = httpClient
	}
}

func WithTimeout(d time.Duration) ClientOption {
	return func(c *Client) {
		if c.HttpClient == nil {
			c.HttpClient = &http.Client{}
		}
		c.HttpClient.Timeout = d
	}
}

func WithLogger(l *slog.Logger) ClientOption {
	return func(c *Client) {
		c.Log = l
	}
}

// NewClient creates a new Paystack API client with the given API key.
func New(apiKey string, opts ...ClientOption) *Client {
	parsedUrl, _ := url.Parse(BASE_URL)
	httpClient := &http.Client{
		Timeout: 5 * time.Second,
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	c := &Client{APIKey: apiKey, HttpClient: httpClient, Log: logger, BaseUrl: parsedUrl}

	for _, opt := range opts {
		opt(c)
	}
	c.Transaction = transaction.New(c)
	c.TransactionSplit = transaction_splits.New(c)
	c.Plan = plans.New(c)
	c.Subscription = subscription.New(c)

	return c
}

func Query(key, value string) types.QueryType {
	return types.QueryType{
		Key:   key,
		Value: value,
	}
}

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

// NewClient creates a new Paystack API client with the given API key.
func New(apiKey string) *Client {
	httpClient := &http.Client{
		Timeout: 5 * time.Second,
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	parsedUrl, _ := url.Parse(BASE_URL)
	c := &Client{APIKey: apiKey, HttpClient: httpClient, Log: logger, BaseUrl: parsedUrl}
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

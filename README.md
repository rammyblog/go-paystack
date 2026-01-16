# go-paystack

A Go client library for [Paystack](https://paystack.com), a payments platform that allows you to accept payments from customers in 154+ countries.

## Installation

```bash
go get github.com/rammyblog/go-paystack
```

## Usage

```go
package main

import (
	"context"
	"log"
	"time"

	"github.com/rammyblog/go-paystack"
	"github.com/rammyblog/go-paystack/transaction"
)

func main() {
	// Create client with default settings
	client := paystack.New("sk_test_your_secret_key")

	// Or with custom options
	client := paystack.New("sk_test_your_secret_key",
		paystack.WithTimeout(10*time.Second),
	)

	ctx := context.Background()

	// Initialize a transaction
	resp, err := client.Transaction.Initialize(ctx, &transaction.TransactionRequest{
		Amount:      100000, // Amount in kobo (1000 NGN)
		Email:       "customer@example.com",
		Currency:    "NGN",
		Reference:   "unique_reference",
		CallbackURL: "https://yoursite.com/callback",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Authorization URL: %s\n", resp.AuthorizationURL)

	// Verify a transaction
	txn, err := client.Transaction.Verify(ctx, "reference")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Transaction status: %s\n", txn.Status)

	// List transactions with query parameters
	list, err := client.Transaction.List(ctx, paystack.Query("page", "1"))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Total transactions: %d\n", list.Meta.Total)
}
```

## Available Services

- **Transaction** - Initialize, verify, list, charge, export transactions
- **Plans** - Create, list, fetch, update subscription plans
- **Subscription** - Create, list, enable, disable subscriptions
- **Transaction Splits** - Create and manage split payments

## Client Options

```go
// Custom HTTP client
paystack.New(apiKey, paystack.WithHTTPClient(customClient))

// Custom timeout
paystack.New(apiKey, paystack.WithTimeout(30*time.Second))

// Custom logger
paystack.New(apiKey, paystack.WithLogger(customLogger))
```

## Examples

See the [example](./example) directory for more usage examples.

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

[MIT](https://choosealicense.com/licenses/mit/)

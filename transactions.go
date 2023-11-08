package paystack

import (
	"context"
)

type Transaction struct {
	client *Client
}

// TransactionRequest represents a request to start a transaction.
type TransactionRequest struct {
	Amount            float32  `json:"amount,omitempty"`
	Email             string   `json:"email,omitempty"`
	Currency          string   `json:"currency,omitempty"`
	Reference         string   `json:"reference,omitempty"`
	CallbackURL       string   `json:"callback_url,omitempty"`
	Plan              string   `json:"plan,omitempty"`
	InvoiceLimit      int      `json:"invoice_limit,omitempty"`
	Metadata          Metadata `json:"metadata,omitempty"`
	Channels          []string `json:"channels,omitempty"`
	SplitCode         string   `json:"split_code,omitempty"`
	SubAccount        string   `json:"subaccount,omitempty"`
	TransactionCharge int      `json:"transaction_charge,omitempty"`
	AuthorizationCode string   `json:"authorization_code,omitempty"`
	Bearer            string   `json:"bearer,omitempty"`
}

type InitializeTransactionResponse struct {
	AuthorizationURL string `json:"authorization_url,omitempty"`
	AccessCode       string `json:"access_code,omitempty"`
	Reference        string `json:"reference,omitempty"`
}

func newTransaction(client *Client) *Transaction {
	return &Transaction{
		client: client,
	}
}

// Initialize initiates a transaction
// For more details see https://paystack.com/docs/api/transaction/#initialize
func (t *Transaction) Initialize(ctx context.Context, txn *TransactionRequest) (*InitializeTransactionResponse, error) {
	url := "transaction/initialize"
	resp := &InitializeTransactionResponse{}
	err := postResource(ctx, t.client, url, txn, resp)
	return resp, err
}

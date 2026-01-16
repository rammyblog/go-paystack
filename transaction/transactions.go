package transaction

import (
	"context"
	"fmt"

	"github.com/rammyblog/go-paystack/internal/helpers"
	"github.com/rammyblog/go-paystack/internal/types"
)

type Transaction struct {
	client types.Requester
}

func New(client types.Requester) *Transaction {
	return &Transaction{
		client: client,
	}
}

// Initialize a transaction
// For more details see https://paystack.com/docs/api/transaction/#initialize
func (t *Transaction) Initialize(ctx context.Context, txn *TransactionRequest) (*InitializeTransactionResponse, error) {
	url := "transaction/initialize"
	resp := &InitializeTransactionResponse{}
	err := t.client.PostResource(ctx, url, txn, resp)
	return resp, err
}

//	Verify a transaction
//
// For more details see https://paystack.com/docs/api/transaction/#verify
func (t *Transaction) Verify(ctx context.Context, reference string) (*TransactionResponse, error) {
	url := fmt.Sprintf("/transaction/verify/%s", reference)
	resp := &TransactionResponse{}
	err := t.client.GetResource(ctx, url, resp)
	return resp, err
}

//	List of transactions
//
// For more details see https://paystack.com/docs/api/transaction/#list
func (t *Transaction) List(ctx context.Context, params ...types.QueryType) (*TransactionList, error) {
	var url string
	if len(params) > 0 {
		url = helpers.AddQueryToUrl("transaction", params...)
	} else {
		url = "/transaction"
	}
	resp := &TransactionList{}

	err := t.client.GetResource(ctx, url, resp)
	return resp, err
}

// Fetch a transaction
// For more details see https://paystack.com/docs/api/transaction/#fetch
func (t *Transaction) FetchById(ctx context.Context, id int) (*TransactionResponse, error) {
	url := fmt.Sprintf("/transaction/%v", id)
	resp := &TransactionResponse{}
	err := t.client.GetResource(ctx, url, resp)
	return resp, err
}

// Charge Authorization
// For more details see https://paystack.com/docs/api/transaction/#charge-authorization
func (t *Transaction) Charge(ctx context.Context, txn *TransactionRequest) (*TransactionResponse, error) {
	url := "transaction/charge_authorization"
	resp := &TransactionResponse{}
	err := t.client.PostResource(ctx, url, txn, resp)
	return resp, err
}

// Total Transactions
// For more details see https://paystack.com/docs/api/transaction/#totals
func (t *Transaction) Total(ctx context.Context, params ...types.QueryType) (*types.GenericResponse, error) {
	var url string
	if len(params) > 0 {
		url = helpers.AddQueryToUrl("transaction/totals", params...)
	} else {
		url = "/transaction/totals"
	}
	resp := &types.GenericResponse{}
	err := t.client.GetResource(ctx, url, resp)
	return resp, err
}

// View the timeline of a transaction
// For more details see https://paystack.com/docs/api/transaction/#view-timeline
func (t *Transaction) Timeline(ctx context.Context, id_or_ref string) (*TransactionResponse, error) {
	url := fmt.Sprintf("/transaction/timeline/%v", id_or_ref)
	resp := &TransactionResponse{}
	err := t.client.GetResource(ctx, url, resp)
	return resp, err
}

// Export your transaction
// For more details see https://paystack.com/docs/api/transaction/#export
func (t *Transaction) Export(ctx context.Context, params ...types.QueryType) (*ExportTransaction, error) {
	var url string
	if len(params) > 0 {
		url = helpers.AddQueryToUrl("transaction/export", params...)
	} else {
		url = "/transaction/export"
	}
	resp := &ExportTransaction{}
	err := t.client.GetResource(ctx, url, resp)
	return resp, err
}

// Partial debit
// For more details see https://paystack.com/docs/api/transaction/#partial-debit
func (t *Transaction) PartialDebit(ctx context.Context, txn *TransactionRequest) (*TransactionResponse, error) {
	url := "transaction/partial_debit"
	resp := &TransactionResponse{}
	err := t.client.PostResource(ctx, url, txn, resp)
	return resp, err
}

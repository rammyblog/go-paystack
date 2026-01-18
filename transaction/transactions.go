package transaction

import (
	"context"
	"fmt"

	"github.com/rammyblog/go-paystack/internal/helpers"
	"github.com/rammyblog/go-paystack/internal/types"
)

type TransactionService interface {
	Initialize(ctx context.Context, txn *TransactionRequest) (*InitializeTransactionResponse, error)
	Verify(ctx context.Context, reference string) (*TransactionResponse, error)
	List(ctx context.Context, params ...types.QueryType) (*TransactionList, error)
	FetchById(ctx context.Context, id int) (*TransactionResponse, error)
	Charge(ctx context.Context, txn *TransactionRequest) (*TransactionResponse, error)
	Total(ctx context.Context, params ...types.QueryType) (*types.GenericResponse, error)
	Timeline(ctx context.Context, id_or_ref string) (*TransactionResponse, error)
	Export(ctx context.Context, params ...types.QueryType) (*ExportTransaction, error)
	PartialDebit(ctx context.Context, txn *TransactionRequest) (*TransactionResponse, error)
}

const (
	INITIALIZE_TRANSACTION    = "transaction/initialize"
	VERIFY_TRANSACTION        = "transaction/verify/%s"
	LIST_TRANSACTION          = "transaction"
	FETCH_TRANSACTION         = "transaction/%d"
	CHARGE_AUTHORIZATION      = "transaction/charge_authorization"
	TOTAL_TRANSACTION         = "transaction/totals"
	TIMELINE_TRANSACTION      = "transaction/timeline/%s"
	EXPORT_TRANSACTION        = "transaction/export"
	PARTIAL_DEBIT_TRANSACTION = "transaction/partial_debit"
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
	url := INITIALIZE_TRANSACTION
	resp := &InitializeTransactionResponse{}
	err := t.client.PostResource(ctx, url, txn, resp)
	return resp, err
}

//	Verify a transaction
//
// For more details see https://paystack.com/docs/api/transaction/#verify
func (t *Transaction) Verify(ctx context.Context, reference string) (*TransactionResponse, error) {
	url := fmt.Sprintf(VERIFY_TRANSACTION, reference)
	resp := &TransactionResponse{}
	err := t.client.GetResource(ctx, url, resp)
	return resp, err
}

//	List of transactions
//
// For more details see https://paystack.com/docs/api/transaction/#list
func (t *Transaction) List(ctx context.Context, params ...types.QueryType) (*TransactionList, error) {
	url := helpers.AddQueryToUrl(LIST_TRANSACTION, params...)
	resp := &TransactionList{}
	err := t.client.GetResource(ctx, url, resp)
	return resp, err
}

// Fetch a transaction
// For more details see https://paystack.com/docs/api/transaction/#fetch
func (t *Transaction) FetchById(ctx context.Context, id int) (*TransactionResponse, error) {
	url := fmt.Sprintf(FETCH_TRANSACTION, id)
	resp := &TransactionResponse{}
	err := t.client.GetResource(ctx, url, resp)
	return resp, err
}

// Charge Authorization
// For more details see https://paystack.com/docs/api/transaction/#charge-authorization
func (t *Transaction) Charge(ctx context.Context, txn *TransactionRequest) (*TransactionResponse, error) {
	url := CHARGE_AUTHORIZATION
	resp := &TransactionResponse{}
	err := t.client.PostResource(ctx, url, txn, resp)
	return resp, err
}

// Total Transactions
// For more details see https://paystack.com/docs/api/transaction/#totals
func (t *Transaction) Total(ctx context.Context, params ...types.QueryType) (*types.GenericResponse, error) {
	url := helpers.AddQueryToUrl(TOTAL_TRANSACTION, params...)
	resp := &types.GenericResponse{}
	err := t.client.GetResource(ctx, url, resp)
	return resp, err
}

// View the timeline of a transaction
// For more details see https://paystack.com/docs/api/transaction/#view-timeline
func (t *Transaction) Timeline(ctx context.Context, id_or_ref string) (*TransactionResponse, error) {
	url := fmt.Sprintf(TIMELINE_TRANSACTION, id_or_ref)
	resp := &TransactionResponse{}
	err := t.client.GetResource(ctx, url, resp)
	return resp, err
}

// Export your transaction
// For more details see https://paystack.com/docs/api/transaction/#export
func (t *Transaction) Export(ctx context.Context, params ...types.QueryType) (*ExportTransaction, error) {
	url := helpers.AddQueryToUrl(EXPORT_TRANSACTION, params...)
	resp := &ExportTransaction{}
	err := t.client.GetResource(ctx, url, resp)
	return resp, err
}

// Partial debit
// For more details see https://paystack.com/docs/api/transaction/#partial-debit
func (t *Transaction) PartialDebit(ctx context.Context, txn *TransactionRequest) (*TransactionResponse, error) {
	url := PARTIAL_DEBIT_TRANSACTION
	resp := &TransactionResponse{}
	err := t.client.PostResource(ctx, url, txn, resp)
	return resp, err
}

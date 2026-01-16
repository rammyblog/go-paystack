package transaction_splits

import (
	"context"
	"fmt"

	"github.com/rammyblog/go-paystack/internal/helpers"
	"github.com/rammyblog/go-paystack/internal/types"
)

type TransactionSplit struct {
	client types.Requester
}

func New(client types.Requester) *TransactionSplit {
	return &TransactionSplit{
		client: client,
	}
}

// Create a transaction split
// For more details see https://paystack.com/docs/api/split/#create
func (ts *TransactionSplit) Create(ctx context.Context, txn *TransactionSplitRequest) (*TransactionSplitResponse, error) {
	url := "split"
	resp := &TransactionSplitResponse{}
	err := ts.client.PostResource(ctx, url, txn, resp)
	return resp, err
}

//	List/Search Transaction splits
//
// For more details see https://paystack.com/docs/api/split/#list
func (ts *TransactionSplit) List(ctx context.Context, params ...types.QueryType) (*TransactionSplitList, error) {
	var url string
	if len(params) > 0 {
		url = helpers.AddQueryToUrl("split", params...)
	} else {
		url = "/split"
	}
	resp := &TransactionSplitList{}

	err := ts.client.GetResource(ctx, url, resp)
	return resp, err
}

//	Fetch a Transaction splits
//
// For more details see https://paystack.com/docs/api/split/#fetch
func (ts *TransactionSplit) Fetch(ctx context.Context, id string) (*TransactionSplitResponse, error) {

	url := fmt.Sprintf("/split/%s", id)

	resp := &TransactionSplitResponse{}

	err := ts.client.GetResource(ctx, url, resp)
	return resp, err
}

// Update a transaction split
// For more details see https://paystack.com/docs/api/split/#update
func (ts *TransactionSplit) Update(ctx context.Context, id string, txn *TransactionSplitRequest) (*TransactionSplitResponse, error) {
	url := fmt.Sprintf("/split/%s", id)
	resp := &TransactionSplitResponse{}
	err := ts.client.PutResource(ctx, url, txn, resp)
	return resp, err
}

// Add/Update a transaction split subaccount
// For more details see https://paystack.com/docs/api/split/#add-subaccount
func (ts *TransactionSplit) AddSubaccount(ctx context.Context, id string, txn *TransactionSplitSubAccount) (*TransactionSplitResponse, error) {
	url := fmt.Sprintf("/split/%s/subaccount/add", id)
	resp := &TransactionSplitResponse{}
	err := ts.client.PostResource(ctx, url, txn, resp)
	return resp, err
}

// Remove a transaction split subaccount
// For more details see https://paystack.com/docs/api/split/#remove-subaccount
func (ts *TransactionSplit) RemoveSubaccount(ctx context.Context, id string, txn *TransactionSplitSubAccount) (*types.GenericResponse, error) {
	url := fmt.Sprintf("/split/%s/subaccount/remove", id)
	resp := &types.GenericResponse{}
	err := ts.client.PostResource(ctx, url, txn, resp)
	return resp, err
}

package paystack

import (
	"context"
	"fmt"
)

type TransactionSplit struct {
	client *Client
}

func newTransactionSplit(client *Client) *TransactionSplit {
	return &TransactionSplit{
		client: client,
	}
}

type TransactionSplitSubAccount struct {
	Subaccount string `json:"subaccount,omitempty"`
	Share      int    `json:"share,omitempty"`
}

type TransactionSplitRequest struct {
	Name             string                       `json:"name,omitempty"`
	Type             string                       `json:"type,omitempty"`
	Currency         string                       `json:"currency,omitempty"`
	Subaccounts      []TransactionSplitSubAccount `json:"subaccounts,omitempty"`
	BearerType       string                       `json:"bearer_type,omitempty"`
	BearerSubaccount string                       `json:"bearer_subaccount,omitempty"`
	Active           bool                         `json:"active,omitempty"`
}

type TransactionSplitList struct {
	Meta PaginationMeta             `json:"meta"`
	Data []TransactionSplitResponse `json:"data"`
}

type TransactionSplitResponse struct {
	ID               int    `json:"id,omitempty"`
	Name             string `json:"name,omitempty"`
	Type             string `json:"type,omitempty"`
	Currency         string `json:"currency,omitempty"`
	Integration      int    `json:"integration,omitempty"`
	Domain           string `json:"domain,omitempty"`
	SplitCode        string `json:"split_code,omitempty"`
	Active           bool   `json:"active,omitempty"`
	BearerType       string `json:"bearer_type,omitempty"`
	BearerSubaccount int    `json:"bearer_subaccount,omitempty"`
	CreatedAt        string `json:"createdAt,omitempty"`
	UpdatedAt        string `json:"updatedAt,omitempty"`
	Subaccounts      []struct {
		Subaccount struct {
			ID                  int         `json:"id,omitempty"`
			SubaccountCode      string      `json:"subaccount_code,omitempty"`
			BusinessName        string      `json:"business_name,omitempty"`
			Description         string      `json:"description,omitempty"`
			PrimaryContactName  interface{} `json:"primary_contact_name,omitempty"`
			PrimaryContactEmail interface{} `json:"primary_contact_email,omitempty"`
			PrimaryContactPhone interface{} `json:"primary_contact_phone,omitempty"`
			Metadata            interface{} `json:"metadata,omitempty"`
			PercentageCharge    int         `json:"percentage_charge,omitempty"`
			SettlementBank      string      `json:"settlement_bank,omitempty"`
			AccountNumber       string      `json:"account_number,omitempty"`
		} `json:"subaccount,omitempty"`
		Share int `json:"share,omitempty"`
	} `json:"subaccounts,omitempty"`
	TotalSubaccounts int `json:"total_subaccounts,omitempty"`
}

// Create a transaction split
// For more details see https://paystack.com/docs/api/split/#create
func (ts *TransactionSplit) Create(ctx context.Context, txn *TransactionSplitRequest) (*TransactionSplitResponse, error) {
	url := "split"
	resp := &TransactionSplitResponse{}
	err := postResource(ctx, ts.client, url, txn, resp)
	return resp, err
}

//	List/Search Transaction splits
//
// For more details see https://paystack.com/docs/api/split/#list
func (ts *TransactionSplit) List(ctx context.Context, params ...QueryType) (*TransactionSplitList, error) {
	var url string
	if len(params) > 0 {
		url = addQueryToUrl("split", params...)
	} else {
		url = "/split"
	}
	resp := &TransactionSplitList{}

	err := getResource(ctx, ts.client, url, resp)
	return resp, err
}

//	Fetch a Transaction splits
//
// For more details see https://paystack.com/docs/api/split/#fetch
func (ts *TransactionSplit) Fetch(ctx context.Context, id string) (*TransactionSplitResponse, error) {

	url := fmt.Sprintf("/split/%s", id)

	resp := &TransactionSplitResponse{}

	err := getResource(ctx, ts.client, url, resp)
	return resp, err
}

// Update a transaction split
// For more details see https://paystack.com/docs/api/split/#update
func (ts *TransactionSplit) Update(ctx context.Context, id string, txn *TransactionSplitRequest) (*TransactionSplitResponse, error) {
	url := fmt.Sprintf("/split/%s", id)
	resp := &TransactionSplitResponse{}
	err := putResource(ctx, ts.client, url, txn, resp)
	return resp, err
}

// Add/Update a transaction split subaccount
// For more details see https://paystack.com/docs/api/split/#add-subaccount
func (ts *TransactionSplit) AddSubaccount(ctx context.Context, id string, txn *TransactionSplitSubAccount) (*TransactionSplitResponse, error) {
	url := fmt.Sprintf("/split/%s/subaccount/add", id)
	resp := &TransactionSplitResponse{}
	err := postResource(ctx, ts.client, url, txn, resp)
	return resp, err
}

// Remove a transaction split subaccount
// For more details see https://paystack.com/docs/api/split/#remove-subaccount
func (ts *TransactionSplit) RemoveSubaccount(ctx context.Context, id string, txn *TransactionSplitSubAccount) (*GenericResponse, error) {
	url := fmt.Sprintf("/split/%s/subaccount/remove", id)
	resp := &GenericResponse{}
	err := postResource(ctx, ts.client, url, txn, resp)
	return resp, err
}

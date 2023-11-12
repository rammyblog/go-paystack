package paystack

import (
	"context"
	"fmt"
)

type Transaction struct {
	client *Client
}

// TransactionList is a list object for transactions.
type TransactionList struct {
	Meta PaginationMeta        `json:"meta"`
	Data []TransactionResponse `json:"data"`
}

type Authorization struct {
	AuthorizationCode string      `json:"authorization_code"`
	Bin               string      `json:"bin"`
	Last4             string      `json:"last4"`
	ExpMonth          string      `json:"exp_month"`
	ExpYear           string      `json:"exp_year"`
	Channel           string      `json:"channel"`
	CardType          string      `json:"card_type"`
	Bank              string      `json:"bank"`
	CountryCode       string      `json:"country_code"`
	Brand             string      `json:"brand"`
	Reusable          bool        `json:"reusable"`
	Signature         string      `json:"signature"`
	AccountName       interface{} `json:"account_name"`
}

type Customer struct {
	ID                       int         `json:"id"`
	FirstName                interface{} `json:"first_name"`
	LastName                 interface{} `json:"last_name"`
	Email                    string      `json:"email"`
	CustomerCode             string      `json:"customer_code"`
	Phone                    interface{} `json:"phone"`
	Metadata                 interface{} `json:"metadata"`
	RiskAction               string      `json:"risk_action"`
	InternationalFormatPhone interface{} `json:"international_format_phone"`
}

type TransactionTimeline struct {
	StartTime      int           `json:"start_time"`
	TimeSpent      int           `json:"time_spent,omitempty"`
	Attempts       int           `json:"attempts,omitempty"`
	Authentication interface{}   `json:"authentication,omitempty"`
	Errors         int           `json:"errors,omitempty"`
	Success        bool          `json:"success,omitempty"`
	Mobile         bool          `json:"mobile,omitempty"`
	Input          []interface{} `json:"input,omitempty"`
	Channel        string        `json:"channel,omitempty"`
	History        []struct {
		Type    string `json:"type,omitempty"`
		Message string `json:"message,omitempty"`
		Time    int    `json:"time,omitempty"`
	} `json:"history,omitempty"`
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
	Queue             bool     `json:"queue,omitempty"`
}

type InitializeTransactionResponse struct {
	AuthorizationURL string `json:"authorization_url,omitempty"`
	AccessCode       string `json:"access_code,omitempty"`
	Reference        string `json:"reference,omitempty"`
}
type TransactionResponse struct {
	ID                 int                 `json:"id"`
	Domain             string              `json:"domain"`
	Status             string              `json:"status"`
	Reference          string              `json:"reference"`
	Amount             int                 `json:"amount"`
	Message            interface{}         `json:"message"`
	GatewayResponse    string              `json:"gateway_response"`
	Channel            string              `json:"channel"`
	Currency           string              `json:"currency"`
	IPAddress          string              `json:"ip_address"`
	Metadata           interface{}         `json:"metadata"`
	Log                TransactionTimeline `json:"log"`
	Fees               int                 `json:"fees"`
	FeesSplit          interface{}         `json:"fees_split"`
	Authorization      Authorization       `json:"authorization"`
	Customer           Customer            `json:"customer"`
	Plan               interface{}         `json:"plan"`
	Split              interface{}         `json:"split"`
	OrderID            interface{}         `json:"order_id"`
	PaidAt             string              `json:"paidAt"`
	CreatedAt          string              `json:"createdAt"`
	RequestedAmount    int                 `json:"requested_amount"`
	PosTransactionData interface{}         `json:"pos_transaction_data"`
	Source             interface{}         `json:"source"`
	FeesBreakdown      interface{}         `json:"fees_breakdown"`
	TransactionDate    string              `json:"transaction_date"`
	PlanObject         interface{}         `json:"plan_object"`
	Subaccount         interface{}         `json:"subaccount"`
}

type ExportTransaction struct {
	Path string `json:"path,omitempty"`
}

func newTransaction(client *Client) *Transaction {
	return &Transaction{
		client: client,
	}
}

// Initialize a transaction
// For more details see https://paystack.com/docs/api/transaction/#initialize
func (t *Transaction) Initialize(ctx context.Context, txn *TransactionRequest) (*InitializeTransactionResponse, error) {
	url := "transaction/initialize"
	resp := &InitializeTransactionResponse{}
	err := postResource(ctx, t.client, url, txn, resp)
	return resp, err
}

//	Verify a transaction
//
// For more details see https://paystack.com/docs/api/transaction/#verify
func (t *Transaction) Verify(ctx context.Context, reference string) (*TransactionResponse, error) {
	url := fmt.Sprintf("/transaction/verify/%s", reference)
	resp := &TransactionResponse{}
	err := getResource(ctx, t.client, url, resp)
	return resp, err
}

//	List of transactions
//
// For more details see https://paystack.com/docs/api/transaction/#list
func (t *Transaction) List(ctx context.Context, params ...QueryType) (*TransactionList, error) {
	var url string
	if len(params) > 0 {
		url = addQueryToUrl("transaction", params...)
	} else {
		url = "/transaction"
	}
	resp := &TransactionList{}

	err := getResource(ctx, t.client, url, resp)
	return resp, err
}

// Fetch a transaction
// For more details see https://paystack.com/docs/api/transaction/#fetch
func (t *Transaction) FetchById(ctx context.Context, id int) (*TransactionResponse, error) {
	url := fmt.Sprintf("/transaction/%v", id)
	resp := &TransactionResponse{}
	err := getResource(ctx, t.client, url, resp)
	return resp, err
}

// Charge Authorization
// For more details see https://paystack.com/docs/api/transaction/#charge-authorization
func (t *Transaction) Charge(ctx context.Context, txn *TransactionRequest) (*TransactionResponse, error) {
	url := "transaction/charge_authorization"
	resp := &TransactionResponse{}
	err := postResource(ctx, t.client, url, txn, resp)
	return resp, err
}

// Total Transactions
// For more details see https://paystack.com/docs/api/transaction/#totals
func (t *Transaction) Total(ctx context.Context, params ...QueryType) (*GenericResponse, error) {
	var url string
	if len(params) > 0 {
		url = addQueryToUrl("transaction/totals", params...)
	} else {
		url = "/transaction/totals"
	}
	resp := &GenericResponse{}
	err := getResource(ctx, t.client, url, resp)
	return resp, err
}

// View the timeline of a transaction
// For more details see https://paystack.com/docs/api/transaction/#view-timeline
func (t *Transaction) Timeline(ctx context.Context, id_or_ref string) (*TransactionResponse, error) {
	url := fmt.Sprintf("/transaction/timeline/%v", id_or_ref)
	resp := &TransactionResponse{}
	err := getResource(ctx, t.client, url, resp)
	return resp, err
}

// Export your transaction
// For more details see https://paystack.com/docs/api/transaction/#export
func (t *Transaction) Export(ctx context.Context, params ...QueryType) (*ExportTransaction, error) {
	var url string
	if len(params) > 0 {
		url = addQueryToUrl("transaction/export", params...)
	} else {
		url = "/transaction/export"
	}
	resp := &ExportTransaction{}
	err := getResource(ctx, t.client, url, resp)
	return resp, err
}

// Partial debit
// For more details see https://paystack.com/docs/api/transaction/#partial-debit
func (t *Transaction) PartialDebit(ctx context.Context, txn *TransactionRequest) (*TransactionResponse, error) {
	url := "transaction/partial_debit"
	resp := &TransactionResponse{}
	err := postResource(ctx, t.client, url, txn, resp)
	return resp, err
}

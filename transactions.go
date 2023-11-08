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

type Log struct {
	StartTime int           `json:"start_time"`
	TimeSpent int           `json:"time_spent"`
	Attempts  int           `json:"attempts"`
	Errors    int           `json:"errors"`
	Success   bool          `json:"success"`
	Mobile    bool          `json:"mobile"`
	Input     []interface{} `json:"input"`
	History   []struct {
		Type    string `json:"type"`
		Message string `json:"message"`
		Time    int    `json:"time"`
	} `json:"history"`
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

type TransactionListQuery struct {
	PerPage    int    `json:"perPage,omitempty"`
	Page       int    `json:"page,omitempty"`
	Customer   int    `json:"customer,omitempty"`
	TerminalId string `json:"terminalid,omitempty"`
	Status     string `json:"status,omitempty"`
	From       string `json:"from,omitempty"`
	To         string `json:"to,omitempty"`
	Amount     int    `json:"amount,omitempty"`
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
type TransactionResponse struct {
	ID                 int                    `json:"id"`
	Domain             string                 `json:"domain"`
	Status             string                 `json:"status"`
	Reference          string                 `json:"reference"`
	Amount             int                    `json:"amount"`
	Message            interface{}            `json:"message"`
	GatewayResponse    string                 `json:"gateway_response"`
	Channel            string                 `json:"channel"`
	Currency           string                 `json:"currency"`
	IPAddress          string                 `json:"ip_address"`
	Metadata           map[string]interface{} `json:"metadata"`
	Log                Log                    `json:"log"`
	Fees               int                    `json:"fees"`
	FeesSplit          interface{}            `json:"fees_split"`
	Authorization      Authorization          `json:"authorization"`
	Customer           Customer               `json:"customer"`
	Plan               interface{}            `json:"plan"`
	Split              interface{}            `json:"split"`
	OrderID            interface{}            `json:"order_id"`
	PaidAt             string                 `json:"paidAt"`
	CreatedAt          string                 `json:"createdAt"`
	RequestedAmount    int                    `json:"requested_amount"`
	PosTransactionData interface{}            `json:"pos_transaction_data"`
	Source             interface{}            `json:"source"`
	FeesBreakdown      interface{}            `json:"fees_breakdown"`
	TransactionDate    string                 `json:"transaction_date"`
	PlanObject         interface{}            `json:"plan_object"`
	Subaccount         interface{}            `json:"subaccount"`
}

func newTransaction(client *Client) *Transaction {
	return &Transaction{
		client: client,
	}
}

func buildQuery(key, value interface{}) string {
	return fmt.Sprintf("%s=%v", key, value)
}

func buildTransactionListQuery(params ...TransactionListQuery) string {
	query := "?"
	for _, param := range params {
		if param.PerPage > 0 {
			query += buildQuery("perPage", param.PerPage)
		}
		if param.Page > 0 {
			query += buildQuery("&page", param.Page)
		}
		if param.Amount > 0 {
			query += buildQuery("&amount", param.Amount)

		}
		if param.From != "" {
			query += buildQuery("&from", param.From)

		}
		if param.Status != "" {
			query += buildQuery("&status", param.Status)

		}
		if param.TerminalId != "" {
			query += buildQuery("&terminalid", param.TerminalId)

		}
		if param.To != "" {
			query += buildQuery("&to", param.To)

		}
		if param.Customer > 0 {
			query += buildQuery("&customer", param.Customer)
		}

	}

	return query
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

func (t *Transaction) ListN(ctx context.Context, params ...TransactionListQuery) (*TransactionList, error) {
	var url string
	if len(params) > 0 {
		url = fmt.Sprintf("/transaction%s", buildTransactionListQuery(params...))
	} else {
		url = "/transaction"
	}
	resp := &TransactionList{}

	err := getResource(ctx, t.client, url, resp)
	return resp, err
}
